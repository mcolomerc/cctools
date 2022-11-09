package services

import (
	"encoding/json"
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/model"
)

type MdsService struct {
	RestClient client.RestClient
	Conf       config.RuntimeConfig
	ClusterUrl string
}

func NewMdsService(conf config.RuntimeConfig) *MdsService {
	restClient := client.New(conf)
	return &MdsService{
		RestClient: *restClient,
		Conf:       conf,
		ClusterUrl: fmt.Sprintf("%s/security/1.0/", conf.UserConfig.EndpointUrl),
	}
}

func (mService *MdsService) Export() {

}

func (mService *MdsService) GetTopicBindings(topic string) ([]model.RoleBinding, error) {
	roles := [3]string{"DeveloperRead", "DeveloperWrite", "ResourceOwner"}
	var topicBindings []model.RoleBinding
	for _, role := range roles {
		roleBindings, err := mService.getTopicBindingsForRole(topic, role)
		if err != nil {
			log.Printf("client: error getting topic role bindings for rol and topic: %s\n", err)
			return nil, err
		}
		topicBindings = append(topicBindings, roleBindings...)
	}
	return topicBindings, nil
}

func (mService *MdsService) getTopicBindingsForRole(topic string, role string) ([]model.RoleBinding, error) {
	log.Printf("topic: %s ", topic)
	log.Printf(mService.ClusterUrl)
	log.Printf("role: %s ", role)
	requestURL := mService.ClusterUrl + "lookup/role/" + role + "/resource/Topic/name/" + topic
    log.Printf("requestUrl: %s", requestURL)


	clusterId := mService.Conf.UserConfig.Cluster

	clusters := &model.Clusters{
		Kafka: clusterId,
	}
	requestBody := &model.Scope{
		Clusters: *clusters,
	}

	request, _ := json.Marshal(requestBody)

	bindings, err := mService.RestClient.Post(requestURL, request)

	if err != nil {
		fmt.Printf("client: error getting topic configs : %s\n", err)
		return nil, err
	}
	var bindingsForTopic []model.RoleBinding
	for _, binding := range bindings {
		bindingForTopic := &model.RoleBinding{
			RoleName: role,
			Users:    binding,
		}
		bindingsForTopic = append(bindingsForTopic, *bindingForTopic)
	}

	return bindingsForTopic, nil
}
