package imprt

import (
	"encoding/json"

	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"os"

	"github.com/mitchellh/mapstructure"
)

type Importer interface {
	Import(path string, typed interface{}) ([]interface{}, error)
}

type ResourceImporter struct {
	Importer
	Config config.Config
}

func (i *ResourceImporter) IteratePath(path string, typed interface{}) ([]interface{}, error) {
	var elements []interface{}
	files, err := os.ReadDir(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, file := range files {
		log.Info("Reading file: " + path + "/" + file.Name())
		element, err := decodeJson(path + "/" + file.Name())
		if err != nil {
			log.Error(err)
			return nil, err
		}

		cfg := &mapstructure.DecoderConfig{
			Metadata: nil,
			Result:   &typed,
			TagName:  "json",
		}
		decoder, _ := mapstructure.NewDecoder(cfg)
		decoder.Decode(element)
		elements = append(elements, element)
	}
	return elements, nil
}

func decodeJson(path string) (map[string]interface{}, error) {
	var result map[string]interface{}

	file, _ := os.Open(path)
	defer file.Close()

	json.NewDecoder(file).Decode(&result)
	return result, nil
}
