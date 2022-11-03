package export

import "mcolomerc/cc-tools/pkg/model"

type Exporter interface {
	ExportTopics(topics []model.Topic, outputPath string) error
	ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error
}
