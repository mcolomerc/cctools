package imprt

import (
	"encoding/json"

	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
	"os"
)

type Importer interface {
	ImportTopics() ([]model.Topic, error)
}

type JSONImporter struct {
	Importer Importer
	Config   config.Config
}

func NewImporter(conf config.Config) (Importer, error) {
	return &JSONImporter{
		Config: conf,
	}, nil
}

func (i *JSONImporter) ImportTopics() ([]model.Topic, error) {
	//Get the JSON file(s) from the source
	path := i.Config.Export.Output + "/" + string(config.ExportTopics) + "/" + string(config.Json)
	//Read the JSON file(s)
	topics, err := iterate(path)
	if err != nil {
		return nil, err
	}
	//For each file under the path
	return topics, nil
}

func iterate(path string) ([]model.Topic, error) {
	var topics []model.Topic
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, file := range files {
		topic := model.Topic{}
		log.Info("Reading file: " + path + "/" + file.Name())
		content, err := os.ReadFile(path + "/" + file.Name())
		if err != nil {
			log.Error(err)
			return nil, err
		}
		err = json.Unmarshal([]byte(content), &topic)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
