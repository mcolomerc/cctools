package imprt

import (
	"encoding/json"
	"reflect"

	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"os"
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
		element := reflect.New(reflect.TypeOf(typed))
		log.Info("Reading file: " + path + "/" + file.Name())
		decodeJson(element, path+"/"+file.Name())
		elements = append(elements, element)
	}
	return elements, nil
}

func decodeJson(v interface{}, path string) interface{} {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(v)
	if err != nil {
		log.Error(err)
	}
	return v
}
