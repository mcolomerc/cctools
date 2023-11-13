package imprt

import (
	"mcolomerc/cc-tools/pkg/config"
	"reflect"
)

type JSONImporter struct {
	ResourceImporter
}

func NewImporter(conf config.Config) (*JSONImporter, error) {
	return &JSONImporter{
		ResourceImporter: ResourceImporter{
			Config: conf,
		},
	}, nil
}

func (i *JSONImporter) Import(path string, typed interface{}) ([]interface{}, error) {
	//Get the JSON file(s) from the source
	path = path + "/" + string(config.Json)
	//Create a new instance of the type
	typedInstance := reflect.New(reflect.TypeOf(typed))

	//Read and parse the JSON file(s)
	elements, err := i.ResourceImporter.IteratePath(path, typedInstance)
	if err != nil {
		return nil, err
	}
	return elements, nil
}
