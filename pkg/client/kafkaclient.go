package client

import (
	"context"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
	"mcolomerc/cc-tools/pkg/util"
	"os"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"golang.org/x/exp/slices"
)

type KafkaAdminClient struct {
	Client kafka.AdminClient
}

func NewKafkaAdminClient(cfg config.Config) (*KafkaAdminClient, error) {
	log.Info("Using KafkaAdmin client configuration: " + cfg.Source.BootstrapServer)
	// Create a new AdminClient.
	config := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Source.BootstrapServer,
	}
	for k, v := range cfg.Source.ClientProps {
		config.SetKey(k, v)
	}

	a, err := kafka.NewAdminClient(config)
	if err != nil {
		log.Error("Failed to create Admin client:")
		log.Error(err)
		return nil, err
	}
	return &KafkaAdminClient{
		Client: *a,
	}, nil
}

func (kadmin *KafkaAdminClient) GetTopics(exclude string) ([]model.Topic, error) {
	defer util.Timer("GetTopics")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	metadata, err := kadmin.Client.GetMetadata(nil, false, 5000)
	if err != nil {
		log.Error("Getting Topics request failed:")
		log.Error(err)
		return nil, err
	}
	log.Info("Exclude: " + exclude)
	topicNames := []string{}
	for name := range metadata.Topics {
		if !strings.Contains(name, exclude) {
			topicNames = append(topicNames, name)
		}
	}
	log.Info("Exporting topics: ")
	log.Info(topicNames)

	// Describe topics
	describeTopicsResult, err := kadmin.Client.DescribeTopics(
		ctx, kafka.NewTopicCollectionOfTopicNames(topicNames),
		kafka.SetAdminOptionIncludeAuthorizedOperations(true))
	if err != nil {
		log.Error("Failed to describe topics: %s\n", err)
		os.Exit(1)
	}
	result := []model.Topic{} // create a slice of Topic
	for _, topic := range describeTopicsResult.TopicDescriptions {
		nTopic := model.Topic{}
		nTopic.Name = topic.Name
		nTopic.Partitions = len(topic.Partitions)
		done := make(chan bool, 2)
		go func() {
			nTopic.Configs, err = kadmin.GetTopicConfigs(topic.Name)
			if err != nil {
				log.Error("Failed to get topic configs: " + topic.Name)
			}
			done <- true
		}()
		go func() {
			nTopic.ACLs, err = kadmin.GetACLs(topic.Name)
			if err != nil {
				log.Error("Failed to get topic ACLs: " + topic.Name)
			}
			done <- true
		}()
		for i := 0; i < 2; i++ {
			<-done
		}
		nTopic.RetentionTime = getElementFromTopicConfigs(nTopic.Configs, "retention.ms")
		nTopic.MinIsr = getElementFromTopicConfigs(nTopic.Configs, "min.insync.replicas")
		nTopic.ReplicationFactor = 3 // TODO: Check
		result = append(result, nTopic)
	}
	return result, nil
}

func getElementFromTopicConfigs(topicConfigs []model.TopicConfig, keyToSearch string) string {
	index := slices.IndexFunc(topicConfigs, func(c model.TopicConfig) bool { return c.Name == keyToSearch })
	return topicConfigs[index].Value.(string)
}

// *
// Get TopicConfig
// *
func (kadmin *KafkaAdminClient) GetTopicConfigs(topicName string) ([]model.TopicConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configQuery := kafka.ConfigResource{
		Type: kafka.ResourceTopic,
		Name: topicName,
	}
	describeConfig, err := kadmin.Client.DescribeConfigs(ctx, []kafka.ConfigResource{configQuery}, kafka.SetAdminRequestTimeout(10*time.Second))
	if err != nil {
		log.Error("Failed to describe topics: %s\n", err)
		os.Exit(1)
	}
	results := []model.TopicConfig{} // create a slice of TopicConfig
	for _, config := range describeConfig[0].Config {
		nConfig := model.TopicConfig{}
		nConfig.Name = config.Name
		nConfig.Value = config.Value
		results = append(results, nConfig)
	}
	return results, nil
}

// * Get Topic ACLs
func (kadmin *KafkaAdminClient) GetACLs(topicName string) ([]model.AclBinding, error) {
	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Describe ACLs on cluster.
	// Set Admin options to wait for the request to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		log.Error("ParseDuration(60s)")
	}

	var resourceType kafka.ResourceType
	var resourcePatternType kafka.ResourcePatternType
	var operation kafka.ACLOperation
	var permissionType kafka.ACLPermissionType

	resourceType, err = kafka.ResourceTypeFromString("topic")
	if err != nil {
		log.Error("Describe ACLs request failed:")
		log.Error(err)
		os.Exit(1)
	}
	resourcePatternType, err = kafka.ResourcePatternTypeFromString("ANY")
	if err != nil {
		log.Error("Describe ACLs request failed:")
		log.Error(err)
		os.Exit(1)
	}

	operation, err = kafka.ACLOperationFromString("ANY")
	if err != nil {
		log.Error("Describe ACLs request failed:")
		log.Error(err)
		os.Exit(1)
	}

	permissionType, err = kafka.ACLPermissionTypeFromString("ANY")
	if err != nil {
		log.Error("Describe ACLs request failed:")
		log.Error(err)
		os.Exit(1)
	}

	filter := kafka.ACLBindingFilter{
		Type:                resourceType,
		Name:                topicName,
		ResourcePatternType: resourcePatternType,
		Operation:           operation,
		PermissionType:      permissionType,
	}

	result, err := kadmin.Client.DescribeACLs(
		ctx,
		filter,
		kafka.SetAdminRequestTimeout(maxDur),
	)
	if err != nil {
		log.Error("Describe ACLs request failed:")
		log.Error(err)
		os.Exit(1)
	}

	// Print results
	if result.Error.Code() != kafka.ErrNoError {
		log.Error("Describe ACLs failed, error code: %s, message: %s\n",
			result.Error.Code(), result.Error.String())
	}
	results := []model.AclBinding{} // create a slice of AclBinding
	for _, acl := range result.ACLBindings {
		nAcl := model.AclBinding{}
		nAcl.ResourceType = acl.Type.String()
		nAcl.ResourceName = acl.Name
		nAcl.Host = acl.Host
		nAcl.Operation = acl.Operation.String()
		nAcl.Permission = acl.PermissionType.String()
		nAcl.PatternType = acl.ResourcePatternType.String()
		nAcl.Principal = acl.Principal
		results = append(results, nAcl)
	}
	return results, nil
}