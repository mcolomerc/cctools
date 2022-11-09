package kafkaexp

import (
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaJsonExporter struct {
	export.JsonExporter
}

func NewKafkaJsonExporter() *KafkaJsonExporter {
	return &KafkaJsonExporter{
		JsonExporter: export.JsonExporter{},
	}
}

func (e KafkaJsonExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	return e.JsonExporter.Export(topics, outputPath+"_topics")
}

func (e KafkaJsonExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.JsonExporter.Export(cgroups, outputPath+"_consumer_groups")
}
