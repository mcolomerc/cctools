package imprt

import (
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"

	"github.com/mitchellh/mapstructure"
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
		Topics:         conf.Import.Source + TOPICS_PATH,
		ConsumerGroups: conf.Import.Source + CGROUPS_PATH,
	}
	destClient, err := client.NewKafkaAdminClient(conf.Destination.Kafka)
	if err != nil {
		log.Error("Error creating Kafka Destination Client : " + err.Error())
		return nil, err
	}
	imp, err := NewImporter(conf)
	if err != nil {
		log.Error("Error creating importer : " + err.Error())
		return nil, err
	}
	kafkaService := &KafkaImport{
		DestClient: *destClient,
		Conf:       conf,
		Paths:      *paths,
		Importer:   imp,
	}

	return kafkaService, nil
}

func (k *KafkaImport) Import() {
	log.Info("Kafka service Importing ...")
	// Import Resources
	for _, v := range k.Conf.Import.Resources {
		if v == config.Topic {
			log.Info("Kafka import topics ...")
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
		typed := model.Topic{}
		cfg := &mapstructure.DecoderConfig{
			Metadata: nil,
			Result:   &typed,
			TagName:  "json",
		}
		decoder, _ := mapstructure.NewDecoder(cfg)
		decoder.Decode(o)
		topics = append(topics, typed)
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
