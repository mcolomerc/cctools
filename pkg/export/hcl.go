package export

type HclExporter struct {
}

func (e HclExporter) GetPath() string {
	return "tfvars"
}

func (e HclExporter) Export(res interface{}, outputPath string) error {
	return nil
}
