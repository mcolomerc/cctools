package exporters

import (
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaJsonExporter struct {
	ParentKafkaExporter
	JsonExporter
}

func NewKafkaJsonExporter() *KafkaJsonExporter {
	return &KafkaJsonExporter{
		JsonExporter: JsonExporter{},
	}
}

func (e KafkaJsonExporter) ExportTopic(topic model.Topic, outputPath string) error {
	return e.JsonExporter.Export(topic, outputPath+"/"+topic.Name)
}

func (e KafkaJsonExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	done := make(chan bool, len(topics))
	for _, topic := range topics {
		go func(topic model.Topic) {
			err := e.ExportTopic(topic, outputPath)
			if err != nil {
				log.Error("Error generating JSON for Topic ..." + topic.Name)
				log.Error(err)
			}
			done <- true
		}(topic)
	}
	for i := 0; i < len(topics); i++ {
		<-done
	}
	close(done)
	return nil
}

func (e KafkaJsonExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.JsonExporter.Export(cgroups, outputPath+"_consumer_groups")
}
