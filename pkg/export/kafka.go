package export

import (
	"fmt"

	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/exporters"

	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/util"
)

type KafkaService struct {
	SourceClient   client.KafkaAdminClient
	DestClient     client.KafkaAdminClient
	Conf           config.Config
	mService       MdsService
	KafkaExporters []exporters.KafkaExporter
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
	// restClient := client.NewRestClient(conf.EndpointUrl, conf.Credentials)
	sourceClient, err := client.NewKafkaAdminClient(conf.Source.Kafka)
	if err != nil {
		log.Error("Error creating Kafka Source Client")
		log.Error(err)
		return nil, err
	}

	// MDS Client
	mService := NewMdsService(conf)

	// Exporters initialization -
	// TODO: Review with the --output flag.
	// --output flag sets conf.Export.Exporters.
	// Only one is supported
	// It should allow more than one.
	var exportersList []exporters.KafkaExporter
	for _, v := range conf.Export.Exporters {
		if v == config.Excel {
			exportersList = append(exportersList, &exporters.KafkaExcelExporter{})
		} else if v == config.Json {
			exportersList = append(exportersList, &exporters.KafkaJsonExporter{})
		} else if v == config.Yaml {
			exportersList = append(exportersList, &exporters.KafkaYamlExporter{})
		} else if v == config.Clink {
			exportersList = append(exportersList, exporters.NewKafkaClinkExporter(conf))
		} else if v == config.Cfk {
			exportersList = append(exportersList, exporters.NewKafkaCfkExporter(conf))
		} else if v == config.Hcl {
			exportersList = append(exportersList, exporters.NewHclKafkaExporter())
		} else {
			log.Info("Kafka Exporter: Unrecognized exporter: %v \n ", v)
		}
	}
	paths := &KafkaPaths{
		Topics:         conf.Export.Output + TOPICS_PATH,
		ConsumerGroups: conf.Export.Output + CGROUPS_PATH,
		Acls:           conf.Export.Output + ACLS_PATH,
	}
	kafkaService := &KafkaService{
		SourceClient:   *sourceClient,
		mService:       *mService,
		Conf:           conf,
		KafkaExporters: exportersList,
		Paths:          *paths,
	}
	// Destination optional
	if conf.Destination.Kafka.BootstrapServer != "" {
		log.Info("Using Destination:")
		log.Info(conf.Destination.BootstrapServer)
		// Destination Client
		destClient, err := client.NewKafkaAdminClient(conf.Destination.Kafka)
		if err != nil {
			log.Error("Error creating Kafka Destination Client : ", err)
			log.Error(err)
			return nil, err
		}
		kafkaService.DestClient = *destClient
	}
	return kafkaService, nil
}

// Manage export
func (kService *KafkaService) Export() {
	log.Debug("Kafka Service Exporting")
	exportExecutors := kService.KafkaExporters
	// var err error
	for _, v := range kService.Conf.Export.Resources {
		if v == config.Topic {
			kService.ExportTopics(exportExecutors)
			log.Info("Topic info successfully exported to " + kService.Paths.Topics)
		}
		if v == config.ConsumerGroup {
			kService.ExportConsumerGroups(exportExecutors)
			log.Info("Consumer Group info successfully exported to " + kService.Paths.ConsumerGroups)
		}
	}
}

// Export Consumer Groups from source cluster to --output format
func (kService *KafkaService) ExportConsumerGroups(exportExecutors []exporters.KafkaExporter) {
	log.Debug("Exporting Consumer Group Info")
	util.BuildPath(kService.Paths.ConsumerGroups)
	cGroupsList, err := kService.SourceClient.GetConsumerGroups()
	if err != nil {
		log.Error("Error getting Consumer Groups :")
		log.Error(err)
	}
	done := make(chan bool, len(exportExecutors))
	for _, expExec := range exportExecutors {
		go func(xporter exporters.KafkaExporter) {
			pth := fmt.Sprintf("%s%s/", kService.Paths.ConsumerGroups, xporter.GetPath())
			log.Debug("Building path :: ", pth)
			util.BuildPath(pth)
			err := xporter.ExportConsumerGroups(cGroupsList, pth)
			if err != nil {
				log.Error("Error exporting Consumer Groups :")
				log.Error(err)
			}
			done <- true
		}(expExec)
	}
	for i := 0; i < len(exportExecutors); i++ {
		<-done
	}
	close(done)
}

// Export Topics from source cluster to --output format
// Export = Get Object collection + call exporters
// Could be more generic
func (kService *KafkaService) ExportTopics(exportExecutors []exporters.KafkaExporter) {
	log.Debug("Exporting Topic Info")
	util.BuildPath(kService.Paths.Topics)
	topicList, err := kService.SourceClient.GetTopics(kService.Conf.Export.Topics.Exclude)
	if err != nil {
		log.Error("Error getting Topics :")
		log.Error(err)
	}
	done := make(chan bool, len(exportExecutors))
	for _, expExec := range exportExecutors {
		go func(xporter exporters.KafkaExporter) {
			pth := fmt.Sprintf("%s%s/", kService.Paths.Topics, xporter.GetPath())
			log.Debug("Building path :: ", pth)
			util.BuildPath(pth)
			err := xporter.ExportTopics(topicList, pth)
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
}
