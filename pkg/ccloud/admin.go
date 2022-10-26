package ccloud

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/config"
	"net/http"
)

type CCloudAdmin struct {
	Cfg        config.Config
	ClusterUrl string
	Bearer     string
}

func New(cfg config.Config) (*CCloudAdmin, error) {
	user := cfg.ApiKey + ":" + cfg.ApiSecret
	bearer := b64.StdEncoding.EncodeToString([]byte(user))
	clusterUrl := fmt.Sprintf("%s/kafka/v3/clusters/%s", cfg.CCloudUrl, cfg.Cluster)
	return &CCloudAdmin{
		Cfg:        cfg,
		Bearer:     bearer,
		ClusterUrl: clusterUrl,
	}, nil
}

// 0.region.provider.confluent.cloud/kafka/v3/clusters/{cluster_id}/consumer-groups' \

func (admin *CCloudAdmin) GetConsumerGroups() ([]string, error) {
	requestURL := admin.ClusterUrl + "/consumer-groups"
	cGroups, err := admin.buildRequest(requestURL)
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

func (admin *CCloudAdmin) GetTopics() ([]Topic, error) {
	/*
		curl -H "Authorization: Basic <BASE64-encoded-key-and-secret>" --request GET --url 'https://<REST-endpoint>/kafka/v3/clusters/<cluster-id>/topics'
	*/
	requestURL := admin.ClusterUrl + "/topics"
	topics, err := admin.buildRequest(requestURL)
	if err != nil {
		fmt.Printf("client: error getting topics : %s\n", err)
		return nil, err
	}

	var topicList []Topic
	for _, value := range topics {
		val := value.(map[string]interface{})
		t := &Topic{
			Name:              val["topic_name"].(string),
			Partitions:        val["partitions_count"],
			ReplicationFactor: val["replication_factor"],
		}

		configs, err := admin.GetTopicConfigs(t.Name)
		if err != nil {
			fmt.Printf("client: error getting topic configs : %s\n", err)
		} else {
			t.Configs = configs
		}
		topicList = append(topicList, *t)
	}
	return topicList, nil
}

func (admin *CCloudAdmin) GetTopicConfigs(topic string) ([]TopicConfig, error) {
	requestURL := admin.ClusterUrl + "/topics/" + topic + "/configs"
	configs, err := admin.buildRequest(requestURL)
	if err != nil {
		fmt.Printf("client: error getting topic configs : %s\n", err)
		return nil, err
	}
	var configsTopic []TopicConfig
	for _, value := range configs {
		val := value.(map[string]interface{})
		t := &TopicConfig{
			Name:  val["name"].(string),
			Value: val["value"],
		}
		configsTopic = append(configsTopic, *t)
	}
	return configsTopic, nil
}

func (admin *CCloudAdmin) buildRequest(requestURL string) ([]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+admin.Bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	// fmt.Printf("\n response code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}

	var result map[string]any
	json.Unmarshal([]byte(resBody), &result)

	return result["data"].([]interface{}), nil
}
