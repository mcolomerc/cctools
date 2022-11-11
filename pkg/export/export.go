package export

type Exporter interface {
	GetPath() string
	Export(res interface{}, outputPath string) error
}
