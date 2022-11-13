package services

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/kafkaexp"
	"mcolomerc/cc-tools/pkg/model"
	"mcolomerc/cc-tools/pkg/util"
	"strings"

	"golang.org/x/exp/slices"
)

type KafkaService struct {
	RestClient     client.RestClient
	Conf           config.Config
	mService       MdsService
	ClusterUrl     string
	KafkaExporters []kafkaexp.KafkaExporter
	Paths          KafkaPaths
}

type KafkaPaths struct {
	Topics         string
	ConsumerGroups string
}

const (
	TOPICS_PATH  = "/topics/"
	CGROUPS_PATH = "/consumer_groups/"
)

func NewKafkaService(conf config.Config) *KafkaService {
	restClient := client.New(conf.EndpointUrl, conf.Credentials)
	mService := NewMdsService(conf)
	var exporters []kafkaexp.KafkaExporter
	for _, v := range conf.Export.Exporters {
		if v == config.Excel {
			exporters = append(exporters, &kafkaexp.KafkaExcelExporter{})
		} else if v == config.Json {
			exporters = append(exporters, &kafkaexp.KafkaJsonExporter{})
		} else if v == config.Yaml {
			exporters = append(exporters, &kafkaexp.KafkaYamlExporter{})
		} else if v == config.Clink {
			exporters = append(exporters, kafkaexp.NewKafkaClinkExporter(conf))
		} else if v == config.Cfk {
			exporters = append(exporters, kafkaexp.NewKafkaCfkExporter(conf))
		} else {
			fmt.Printf("Kafka Exporter: Unrecognized exporter: %v \n ", v)
		}
	}
	paths := &KafkaPaths{
		Topics:         conf.Export.Output + TOPICS_PATH,
		ConsumerGroups: conf.Export.Output + CGROUPS_PATH,
	}
	return &KafkaService{
		RestClient:     *restClient,
		mService:       *mService,
		Conf:           conf,
		ClusterUrl:     fmt.Sprintf("%s/kafka/v3/clusters/%s", conf.EndpointUrl, conf.Cluster),
		KafkaExporters: exporters,
		Paths:          *paths,
	}
}

func (kService *KafkaService) Export() {
	var result model.KafkaServiceResult
	exportExecutors := kService.KafkaExporters
	// var err error
	for _, v := range kService.Conf.Export.Resources {
		if v == config.ExportTopics {
			log.Println("Exporting Topic Info")
			util.BuildPath(kService.Paths.Topics)
			result.Topics = kService.GetTopics()
			done := make(chan bool, len(exportExecutors))
			for _, v := range exportExecutors {
				go func(v kafkaexp.KafkaExporter) {
					pth := fmt.Sprintf("%s%s/", kService.Paths.Topics, v.GetPath())
					util.BuildPath(pth)
					err := v.ExportTopics(result.Topics, pth)
					if err != nil {
						fmt.Printf("Error: %s\n", err)
					}
					done <- true
				}(v)
			}
			for i := 0; i < len(exportExecutors); i++ {
				<-done
			}
			close(done)
			log.Println("Topic info successfully exported")
		}
		if v == config.ExportConsumerGroups {
			log.Println("Exporting Consumer Group Info")
			util.BuildPath(kService.Paths.ConsumerGroups)
			cgroups := kService.GetConsumerGroups()
			done := make(chan bool, len(exportExecutors))
			for _, v := range exportExecutors {
				go func(v kafkaexp.KafkaExporter) {
					pth := fmt.Sprintf("%s%s/", kService.Paths.ConsumerGroups, v.GetPath())
					util.BuildPath(pth)
					err := v.ExportConsumerGroups(cgroups, pth)
					if err != nil {
						fmt.Printf("Error: %s\n", err)
					}
					done <- true
				}(v)
			}
			for i := 0; i < len(exportExecutors); i++ {
				<-done
			}
			close(done)
			log.Println("Consumer Group Info successfully exported")
		}
	}
}

func (kService *KafkaService) GetTopics() []model.Topic {
	topics, err := kService.RestClient.GetList(kService.ClusterUrl + "/topics")
	if err != nil {
		log.Printf("Error getting Topics : %s\n", err)
		return nil
	}
	var topicList []model.Topic
	finalTopics := kService.TopicsExclusion(topics)
	done := make(chan model.Topic, len(finalTopics))
	for _, value := range finalTopics {
		go func(t model.Topic) {
			configs, err := kService.GetTopicConfigs(t.Name)
			if err != nil {
				log.Printf("Error getting Topic configs : %s\n", err)
			} else {
				t.RetentionTime = getElementFromTopicConfigs(configs, "retention.ms")
				t.MinIsr = getElementFromTopicConfigs(configs, "min.insync.replicas")
				t.Configs = configs
			}

			if kService.mService.Conf.CCloud.Environment == "" {
				bindings, err := kService.mService.GetResourceBindings("Topic", t.Name)
				if err != nil {
					log.Printf("client: error getting topic role bindings for rol and topic: %s\n", err)
				}
				t.RoleBindings = bindings
			}

			done <- t
		}(value)
	}
	for i := 0; i < len(finalTopics); i++ {
		topicList = append(topicList, <-done)
	}
	return topicList
}

func (kService *KafkaService) TopicsExclusion(topics []interface{}) []model.Topic {
	var finalTopics []model.Topic
	for _, value := range topics {
		val := value.(map[string]interface{})
		topicName := val["topic_name"].(string)
		if !kService.checkExclude(topicName) {
			t := &model.Topic{
				Name:              topicName,
				Partitions:        val["partitions_count"],
				ReplicationFactor: val["replication_factor"],
			}
			finalTopics = append(finalTopics, *t)
		}
	}
	return finalTopics
}

func (kService *KafkaService) GetTopicConfigs(topic string) ([]model.TopicConfig, error) {
	configs, err := kService.RestClient.GetList(kService.ClusterUrl + "/topics/" + topic + "/configs")
	if err != nil {
		log.Printf("Error getting Topic configs : %s\n", err)
		return nil, err
	}
	var configsTopic []model.TopicConfig
	for _, value := range configs {
		val := value.(map[string]interface{})
		t := &model.TopicConfig{
			Name:  val["name"].(string),
			Value: val["value"],
		}
		configsTopic = append(configsTopic, *t)
	}
	return configsTopic, nil
}

func getElementFromTopicConfigs(topicConfigs []model.TopicConfig, keyToSearch string) string {
	index := slices.IndexFunc(topicConfigs, func(c model.TopicConfig) bool { return c.Name == keyToSearch })

	return topicConfigs[index].Value.(string)
}

func (kService *KafkaService) GetConsumerGroups() []model.ConsumerGroup {
	cGroups, err := kService.RestClient.GetList(kService.ClusterUrl + "/consumer-groups")
	if err != nil {
		fmt.Printf("Error getting consumer-groups : %s\n", err)
		return nil
	}

	var consumerGroups []model.ConsumerGroup
	for _, value := range cGroups {
		val := value.(map[string]interface{})
		cg := &model.ConsumerGroup{
			ConsumerGroupID:    val["consumer_group_id"].(string),
			PartitionsAssignor: val["partition_assignor"].(string),
			State:              val["state"].(string),
		}
		consumers := kService.GetConsumers(cg.ConsumerGroupID)
		if err != nil {
			log.Printf("client: error getting consumers for consumer group : %s\n", err)
		} else {
			cg.Consumers = consumers
		}
		cg.LagSummary = kService.GetLagSummary(cg.ConsumerGroupID)
		cg.Lags = kService.GetLag(cg.ConsumerGroupID)
		consumerGroups = append(consumerGroups, *cg)
	}
	return consumerGroups
}
func (kService *KafkaService) GetLag(group string) []model.Lag {
	lagResp, err := kService.RestClient.GetList(kService.ClusterUrl + "/consumer-groups/" + group + "/lags")
	if err != nil {
		fmt.Printf("Error getting consumer groups lags : %s\n", err)
		return nil
	}
	var lags []model.Lag
	for _, value := range lagResp {
		val := value.(map[string]interface{})
		l := &model.Lag{
			TopicName:     val["topic_name"].(string),
			Partition:     val["partition_id"],
			ClientId:      val["client_id"].(string),
			CurrentOffset: val["current_offset"],
			LogEndOffset:  val["log_end_offset"],
			Lag:           val["lag"],
		}
		lags = append(lags, *l)
	}
	return lags

}
func (kService *KafkaService) GetLagSummary(group string) model.LagSummary {
	lagResp, err := kService.RestClient.Get(kService.ClusterUrl + "/consumer-groups/" + group + "/lag-summary")
	fmt.Println(lagResp)
	var lags model.LagSummary
	if err != nil {
		fmt.Printf("Error getting consumers : %s\n", err)
	}
	return lags
}

func (kService *KafkaService) GetConsumers(group string) []model.Consumer {
	consumersResp, err := kService.RestClient.GetList(kService.ClusterUrl + "/consumer-groups/" + group + "/consumers")
	if err != nil {
		fmt.Printf("Error getting consumers : %s\n", err)
		return nil
	}

	var consumers []model.Consumer
	for _, value := range consumersResp {
		val := value.(map[string]interface{})
		var instance string
		if val["instance_id"] == nil {
			instance = ""
		}
		c := &model.Consumer{
			ConsumerId: val["consumer_id"].(string),
			InstanceId: instance,
			ClientId:   val["client_id"].(string),
		}
		consumers = append(consumers, *c)
	}
	return consumers
}

func (kService *KafkaService) checkExclude(topic string) bool {
	if kService.Conf.Export.Topics.Exclude != "" {
		if strings.Contains(topic, kService.Conf.Export.Topics.Exclude) {
			if kService.Conf.Export.Topics.Include != "" {
				if strings.Contains(topic, kService.Conf.Export.Topics.Include) {
					return false
				}
			}
			return true
		}
	} else if kService.Conf.Export.Topics.Include != "" {
		if strings.Contains(topic, kService.Conf.Export.Topics.Include) {
			return false
		} else {
			return true
		}
	}
	return false
}
