package export

import (
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/ccloud"
	"mcolomerc/cc-tools/pkg/config"

	"gopkg.in/yaml.v2"
)

type YamlExporter struct{}

func (e YamlExporter) ExportTopics(topics []ccloud.Topic, config config.Config) error {

	yamlData, err := yaml.Marshal(&topics)

	if err != nil {
		return err
	}

	fileName := config.Export.Output + "/" + config.Cluster + "_Topics.yaml"
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}
