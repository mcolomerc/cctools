package export

type Exporter interface {
	Export(res interface{}, outputPath string) error
}
