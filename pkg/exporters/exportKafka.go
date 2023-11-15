package exporters

import (
	"mcolomerc/cc-tools/pkg/model"
)

type KafkaExporter interface {
	Exporter
	ExportTopics(topics []model.Topic, outputPath string) error
	ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error
}

type ParentKafkaExporter struct {
}

// Abstracts the KafkaExporter interface

func (e ParentKafkaExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	return nil
}
func (e ParentKafkaExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return nil
}
