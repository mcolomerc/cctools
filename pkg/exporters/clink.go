package exporters

type ClinkExporter struct {
	LinkName             string
	SourceClusterId      string
	BootstrapServer      string
	SourceApiKey         string
	SourceApiSecret      string
	DestinationClusterId string
	AclSync              bool
	OffsetSync           bool
	AutoCreate           bool
}

func (e ClinkExporter) GetPath() string {
	return "clink"
}

func (e ClinkExporter) Export(res interface{}, outputPath string) error {
	return nil
}
