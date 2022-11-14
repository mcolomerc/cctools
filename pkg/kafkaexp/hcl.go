package kafkaexp

import (
	"fmt"
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
)

type HclExporter struct {
	ParentKafkaExporter
	export.HclExporter
}

func NewHclExporter() *HclExporter {
	return &HclExporter{
		HclExporter: export.HclExporter{},
	}
}

func (e HclExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	log.Debug("Exporting topics to HCL")

	for i, topic := range topics {
		hclFile := fmt.Sprintf("%s%s%d", outputPath, "topic", i)
		e.HclExporter.Export(topic, hclFile)
	}
	return nil
}

func (e HclExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.HclExporter.Export(cgroups, outputPath+"_consumer_groups")
}
