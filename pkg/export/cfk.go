package export

type CfkExporter struct {
	Namespace      string
	KafkaRestClass string
}

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

func (e CfkExporter) GetPath() string {
	return "cfk"
}

func (e CfkExporter) Export(res interface{}, outputPath string) error {
	return nil
}
