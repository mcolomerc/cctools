package kafka

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaService struct {
	KafkaRestClient client.KafkaRestClient
	Conf            config.RuntimeConfig
	ClusterUrl      string
}

func New(conf config.RuntimeConfig) *KafkaService {
	restClient := client.New(conf)

	return &KafkaService{
		KafkaRestClient: *restClient,
		Conf:            conf,
		ClusterUrl:      fmt.Sprintf("%s/kafka/v3/clusters/%s", conf.UserConfig.EndpointUrl, conf.UserConfig.Cluster),
	}
}

func (kService *KafkaService) GetTopics() ([]model.Topic, error) {
	topics, err := kService.KafkaRestClient.Get(kService.ClusterUrl + "/topics")
	if err != nil {
		log.Printf("client: error getting topics : %s\n", err)
		return nil, err
	}
	var topicList []model.Topic
	for _, value := range topics {
		val := value.(map[string]interface{})
		t := &model.Topic{
			Name:              val["topic_name"].(string),
			Partitions:        val["partitions_count"],
			ReplicationFactor: val["replication_factor"],
		}

		configs, err := kService.GetTopicConfigs(t.Name)
		if err != nil {
			log.Printf("client: error getting topic configs : %s\n", err)
		} else {
			t.Configs = configs
		}
		topicList = append(topicList, *t)
	}
	return topicList, nil
}

func (kService *KafkaService) GetTopicConfigs(topic string) ([]model.TopicConfig, error) {
	configs, err := kService.KafkaRestClient.Get(kService.ClusterUrl + "/topics/" + topic + "/configs")
	if err != nil {
		fmt.Printf("client: error getting topic configs : %s\n", err)
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

func (kService *KafkaService) GetConsumerGroups() ([]string, error) {
	cGroups, err := kService.KafkaRestClient.Get(kService.ClusterUrl + "/consumer-groups")
	if err != nil {
		fmt.Printf("client: error getting consumer-groups : %s\n", err)
		return nil, err
	}

	var consumerGroups []string
	for _, value := range cGroups {
		val := value.(map[string]interface{})
		consumerGroups = append(consumerGroups, val["consumer_group_id"].(string))
	}
	return consumerGroups, nil
}

/*curl --request GET \
  --url 'https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags' \
  --header 'Authorization: Basic REPLACE_BASIC_AUTH'

 "https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1/configs"
},

https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1/configs

https://pkc-00000.region.provider.confluent.cloud/kafka/v3/clusters/cluster-1/topics/topic-1/configs

*/
