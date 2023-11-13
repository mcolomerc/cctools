package exporters

type ExcelExporter struct{}

func (e ExcelExporter) GetPath() string {
	return "xls"
}

func (e ExcelExporter) Export(res interface{}, outputPath string) error {
	return nil
}
