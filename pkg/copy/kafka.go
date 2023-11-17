package copy

import (
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
)

type KafkaCopy struct {
	SourceClient client.KafkaAdminClient
	DestClient   client.KafkaAdminClient
	Conf         config.Config
	Paths        KafkaPaths
}

type KafkaPaths struct {
	Topics         string
	ConsumerGroups string
}

const (
	TOPICS_PATH  = "/topics/"
	CGROUPS_PATH = "/consumer_groups/"
)

func NewKafkaCopy(conf config.Config) (*KafkaCopy, error) {
	paths := &KafkaPaths{
		Topics:         conf.Export.Output + TOPICS_PATH,
		ConsumerGroups: conf.Export.Output + CGROUPS_PATH,
	}
	sourceClient, err := client.NewKafkaAdminClient(conf.Source.Kafka)
	if err != nil {
		log.Error("Error creating Kafka Destination Client : ", err)
		log.Error(err)
		return nil, err
	}
	destClient, err := client.NewKafkaAdminClient(conf.Destination.Kafka)
	if err != nil {
		log.Error("Error creating Kafka Destination Client : ", err)
		log.Error(err)
		return nil, err
	}
	kafkaService := &KafkaCopy{
		SourceClient: *sourceClient,
		DestClient:   *destClient,
		Conf:         conf,
		Paths:        *paths,
	}
	return kafkaService, nil
}

func (k *KafkaCopy) Copy() {
	// Copy Topics
	k.CopyTopics()
	// Copy Consumer Groups
	// k.copyConsumerGroups()
}

func (k *KafkaCopy) CopyTopics() {
	// Get Topics from Source
	topics, err := k.SourceClient.GetTopics(k.Conf.Export.Topics.Exclude)
	if err != nil {
		log.Error("Error getting topics from source cluster: ", err)
		return
	}
	if len(topics) > 0 {
		err = k.DestClient.CreateTopics(topics)
		if err != nil {
			log.Error("Error writing topics to destination cluster: ", err)
			return
		}
	}
}
