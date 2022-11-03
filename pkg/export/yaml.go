package export

import (
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/model"

	"gopkg.in/yaml.v2"
)

type YamlExporter struct{}

func (e YamlExporter) Export(res interface{}, outputPath string) error {
	file, errJson := yaml.Marshal(res)
	if errJson != nil {
		return errJson
	}
	err := ioutil.WriteFile(outputPath+"_.yml", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (e YamlExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	return e.Export(topics, outputPath+"_topics")
}

func (e YamlExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.Export(cgroups, outputPath+"_consumer_groups")
}
