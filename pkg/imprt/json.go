package imprt

import (
	"mcolomerc/cc-tools/pkg/config"
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
	path = path + string(config.Json)
	//Read and parse the JSON file(s)
	elements, err := i.ResourceImporter.IteratePath(path, typed)
	if err != nil {
		return nil, err
	}
	return elements, nil
}
