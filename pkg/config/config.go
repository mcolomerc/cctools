package config

type Config struct {
	Cluster         string `yaml:"cluster" validate:"required"`
	BootstrapServer string `yaml:"bootstrapServer" validate:"required"`
	Environment     string `yaml:"environment" validate:"required"`
	ApiKey          string `yaml:"apiKey" validate:"required"`
	ApiSecret       string `yaml:"apiSecret" validate:"required"`
	CCloudUrl       string `yaml:"ccloudUrl" validate:"required"`
	Export          Export `yaml:"export"`
}
type Export struct {
	Exporters []Exporter `yaml:"exporters"`
	Output    string     `yaml:"output"`
}

//
type Exporter string

const (
	Excel Exporter = "excel"
	Yaml  Exporter = "yaml"
	Json  Exporter = "json"
)
