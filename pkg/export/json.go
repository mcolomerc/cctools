package export

import (
	"encoding/json"
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/ccloud"
	"mcolomerc/cc-tools/pkg/config"
)

type JsonExporter struct{}

func (e JsonExporter) ExportTopics(topics []ccloud.Topic, config config.Config) error {
	file, errJson := json.MarshalIndent(topics, "", " ")
	if errJson != nil {
		return errJson
	}
	err := ioutil.WriteFile(config.Export.Output+"/"+config.Cluster+"_Topics.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
