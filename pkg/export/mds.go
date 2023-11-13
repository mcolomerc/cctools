package export

import (
	"encoding/json"
	"fmt"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
)

type MdsService struct {
	RestClient client.RestClient
	Conf       config.Config
	ClusterUrl string
}

func NewMdsService(conf config.Config) *MdsService {
	restClient := client.NewRestClient(conf.Source.EndpointUrl, conf.Source.Credentials)
	return &MdsService{
		RestClient: *restClient,
		Conf:       conf,
		ClusterUrl: fmt.Sprintf("%s/security/1.0/", conf.Source.EndpointUrl),
	}
}

func (mService *MdsService) Export() {

}

func (mService *MdsService) GetResourceBindings(resourceType string, resourceName string) ([]model.RoleBinding, error) {
	roles := [3]string{"DeveloperRead", "DeveloperWrite", "ResourceOwner"}
	var topicBindings []model.RoleBinding
	for _, role := range roles {
		roleBindings, err := mService.getResourceBindingsForRole(resourceType, resourceName, role)
		if err != nil {
			log.Error("client: error getting topic role bindings for rol and topic: %s\n", err)
			return nil, err
		}
		topicBindings = append(topicBindings, roleBindings...)
	}
	return topicBindings, nil
}

func (mService *MdsService) getResourceBindingsForRole(resourceType string, resourceName string, role string) ([]model.RoleBinding, error) {
	requestURL := mService.ClusterUrl + "lookup/role/" + role +
		"/resource/" + resourceType + "/name/" + resourceName

	clusterId := mService.Conf.Cluster

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
