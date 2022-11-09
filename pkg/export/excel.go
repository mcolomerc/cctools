package export

type ExcelExporter struct{}

func (e ExcelExporter) Export(res interface{}, outputPath string) error {
	return nil
}
