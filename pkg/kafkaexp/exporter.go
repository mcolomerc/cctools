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
