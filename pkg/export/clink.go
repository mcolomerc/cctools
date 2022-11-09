package export

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

const (
	SH_HEADER = "#!/bin/bash"
)

func (e ClinkExporter) Export(res interface{}, outputPath string) error {
	return nil
}
