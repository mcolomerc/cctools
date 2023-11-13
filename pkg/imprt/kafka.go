package imprt

import (
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaImport struct {
	DestClient client.KafkaAdminClient
	Conf       config.Config
	Paths      KafkaPaths
	Importer   Importer
}

type KafkaPaths struct {
	Topics         string
	ConsumerGroups string
}

const (
	TOPICS_PATH  = "/topics/"
	CGROUPS_PATH = "/consumer_groups/"
)

func NewKafkaImport(conf config.Config) (*KafkaImport, error) {
	paths := &KafkaPaths{
		Topics:         conf.Export.Output + TOPICS_PATH,
		ConsumerGroups: conf.Export.Output + CGROUPS_PATH,
	}
	destClient, err := client.NewKafkaAdminClient(conf.Destination.Kafka)
	if err != nil {
		log.Error("Error creating Kafka Destination Client : ", err)
		log.Error(err)
		return nil, err
	}
	imp, err := NewImporter(conf)

	kafkaService := &KafkaImport{
		DestClient: *destClient,
		Conf:       conf,
		Paths:      *paths,
		Importer:   imp,
	}

	return kafkaService, nil
}

func (k *KafkaImport) Import() {
	// Import Resources
	for _, v := range k.Conf.Import.Resources {
		if v == config.Topic {
			k.ImportTopics()
		}
	}
	//k.ImportConsumerGroups()
}

func (i *KafkaImport) ImportTopics() {
	// Import Topics
	objs, err := i.Importer.Import(i.Paths.Topics, model.Topic{})
	if err != nil {
		log.Error("Error importing topics ")
		log.Error(err)
	}
	var topics []model.Topic
	for _, o := range objs {
		topics = append(topics, o.(model.Topic))
	}
	i.WriteTopics(topics)
}

func (i *KafkaImport) WriteTopics(topics []model.Topic) {
	// Create Topics with configs destination cluster
	err := i.DestClient.CreateTopics(topics)
	if err != nil {
		log.Error("Error writing topics ")
		log.Error(err)
	}

	// Create ACLs for each topic
	for _, topic := range topics {
		err = i.DestClient.SetACLs(topic.ACLs, i.Conf.Principals)
		if err != nil {
			log.Error("Error writing ACLs ")
			log.Error(err)
		}
	}
}
