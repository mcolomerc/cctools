package kafkaexp

import (
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaExporter interface {
	export.Exporter
	ExportTopics(topics []model.Topic, outputPath string) error
	ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error
}

type ParentKafkaExporter struct {
}

func (e ParentKafkaExporter) ExportTopics(topics []model.Topic, outputPath string) error { return nil }
func (e ParentKafkaExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return nil
}
