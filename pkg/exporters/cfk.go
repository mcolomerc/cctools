package exporters

type CfkExporter struct {
	Namespace  string
	ClusterRef string
}

func (e CfkExporter) GetPath() string {
	return "cfk"
}

func (e CfkExporter) Export(res interface{}, outputPath string) error {
	return nil
}
