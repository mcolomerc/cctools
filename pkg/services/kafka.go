package services

import (
	"fmt"

	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/kafkaexp"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
	"mcolomerc/cc-tools/pkg/util"
	"strings"
)

type KafkaService struct {
	RestClient     client.RestClient
	AdminClient    client.KafkaAdminClient
	Conf           config.Config
	mService       MdsService
	ClusterUrl     string
	KafkaExporters []kafkaexp.KafkaExporter
	Paths          KafkaPaths
}

type KafkaPaths struct {
	Topics         string
	ConsumerGroups string
	Acls           string
}

const (
	TOPICS_PATH  = "/topics/"
	CGROUPS_PATH = "/consumer_groups/"
	ACLS_PATH    = "/acls/"
)

func NewKafkaService(conf config.Config) (*KafkaService, error) {
	// Kafka Clients initialization
	// REST Client
	restClient := client.NewRestClient(conf.EndpointUrl, conf.Credentials)
	// Admin Client
	adminClient, err := client.NewKafkaAdminClient(conf)
	if err != nil {
		log.Error("Error creating Kafka Admin Client : %s\n", err)
		return nil, err
	}
	// MDS Client
	mService := NewMdsService(conf)

	// Exporters initialization
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
		} else if v == config.Hcl {
			exporters = append(exporters, &kafkaexp.HclExporter{})
		} else {
			log.Info("Kafka Exporter: Unrecognized exporter: %v \n ", v)
		}
	}
	paths := &KafkaPaths{
		Topics:         conf.Export.Output + TOPICS_PATH,
		ConsumerGroups: conf.Export.Output + CGROUPS_PATH,
		Acls:           conf.Export.Output + ACLS_PATH,
	}
	return &KafkaService{
		RestClient:     *restClient,
		AdminClient:    *adminClient,
		mService:       *mService,
		Conf:           conf,
		ClusterUrl:     fmt.Sprintf("%s/kafka/v3/clusters/%s", conf.EndpointUrl, conf.Cluster),
		KafkaExporters: exporters,
		Paths:          *paths,
	}, nil
}

func (kService *KafkaService) Export() {
	log.Debug("Kafka Service Exporting")
	var result model.KafkaServiceResult
	exportExecutors := kService.KafkaExporters
	// var err error
	for _, v := range kService.Conf.Export.Resources {
		if v == config.ExportTopics {
			log.Debug("Exporting Topic Info")
			util.BuildPath(kService.Paths.Topics)
			topicList, err := kService.AdminClient.GetTopics(kService.Conf.Export.Topics.Exclude)
			if err != nil {
				log.Error("Error getting Topics :")
				log.Error(err)
			}
			result.Topics = topicList
			done := make(chan bool, len(exportExecutors))
			for _, expExec := range exportExecutors {
				go func(xporter kafkaexp.KafkaExporter) {
					pth := fmt.Sprintf("%s%s/", kService.Paths.Topics, xporter.GetPath())
					log.Debug("Building path :: ", pth)
					util.BuildPath(pth)
					err := xporter.ExportTopics(result.Topics, pth)
					if err != nil {
						log.Error("Error exporting Topics :")
						log.Error(err)
					}
					done <- true
				}(expExec)
			}
			for i := 0; i < len(exportExecutors); i++ {
				<-done
			}
			close(done)
			log.Info("Topic info successfully exported to " + kService.Paths.Topics)
		}
	}
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
			log.Error("client: error getting consumers for consumer group : %s\n", err)
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
