package export

import (
	"encoding/json"
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/model"
)

type JsonExporter struct{}

func (e JsonExporter) Export(res interface{}, outputPath string) error {
	file, errJson := json.MarshalIndent(res, "", " ")
	if errJson != nil {
		return errJson
	}
	err := ioutil.WriteFile(outputPath+"_.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (e JsonExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	return e.Export(topics, outputPath+"_topics")
}

func (e JsonExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return e.Export(cgroups, outputPath+"_consumer_groups")
}
