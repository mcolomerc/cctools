package export

import (
	"encoding/json"
	"io/ioutil"
)

type JsonExporter struct{}

func (e JsonExporter) GetPath() string {
	return "json"
}

func (e JsonExporter) Export(res interface{}, outputPath string) error {
	file, errJson := json.MarshalIndent(res, "", " ")
	if errJson != nil {
		return errJson
	}
	err := ioutil.WriteFile(outputPath+".json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
