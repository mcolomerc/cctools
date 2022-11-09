package export

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YamlExporter struct{}

func (e YamlExporter) Export(res interface{}, outputPath string) error {
	file, errJson := yaml.Marshal(res)
	if errJson != nil {
		return errJson
	}
	err := ioutil.WriteFile(outputPath+"_.yml", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
