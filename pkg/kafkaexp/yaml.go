package kafkaexp

import (
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaYamlExporter struct {
	export.YamlExporter
}

func NewKafkaYamlExporter() *KafkaYamlExporter {
	return &KafkaYamlExporter{
		YamlExporter: export.YamlExporter{},
	}
}

func (e KafkaYamlExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	return e.YamlExporter.Export(topics, outputPath+"_topics")
}

func (e KafkaYamlExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.YamlExporter.Export(cgroups, outputPath+"_consumer_groups")
}
