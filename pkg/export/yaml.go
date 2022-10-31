package export

import (
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/model"

	"gopkg.in/yaml.v2"
)

type YamlExporter struct{}

func (e YamlExporter) ExportTopics(topics []model.Topic, outputPath string) error {

	yamlData, err := yaml.Marshal(&topics)
	if err != nil {
		return err
	}

	fileName := outputPath + "_Topics.yaml"
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}
