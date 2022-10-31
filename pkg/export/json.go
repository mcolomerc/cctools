package export

import (
	"encoding/json"
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/model"
)

type JsonExporter struct{}

func (e JsonExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	file, errJson := json.MarshalIndent(topics, "", " ")
	if errJson != nil {
		return errJson
	}
	err := ioutil.WriteFile(outputPath+"_Topics.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
