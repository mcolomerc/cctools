package exporters

import (
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaYamlExporter struct {
	ParentKafkaExporter
	YamlExporter
}

func NewKafkaYamlExporter() *KafkaYamlExporter {
	return &KafkaYamlExporter{
		YamlExporter: YamlExporter{},
	}
}

func (e KafkaYamlExporter) ExportTopic(topic model.Topic, outputPath string) error {
	return e.YamlExporter.Export(topic, outputPath+"/"+topic.Name)
}

func (e KafkaYamlExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	done := make(chan bool, len(topics))
	for _, topic := range topics {
		go func(topic model.Topic) {
			err := e.ExportTopic(topic, outputPath)
			if err != nil {
				log.Error("Error generating YAML for Topic ..." + topic.Name)
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

func (e KafkaYamlExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.YamlExporter.Export(cgroups, outputPath+"_consumer_groups")
}
