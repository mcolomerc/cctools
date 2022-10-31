package config

type Config struct {
	Cluster         string      `yaml:"cluster" validate:"required"`
	EndpointUrl     string      `yaml:"endpointUrl" validate:"required"`
	BootstrapServer string      `yaml:"bootstrapServer" validate:"required"`
	Credentials     Credentials `yaml:"credentials"`
	CCloud          CCloud      `yaml:"ccloud" validate:"omitempty"`
	Export          Export      `yaml:"export"`
}

type Credentials struct {
	Key          string       `yaml:"key" validate:"required"`
	Secret       string       `yaml:"secret" validate:"required"`
	Certificates Certificates `yaml:"certificates" validate:"omitempty"`
}

type Certificates struct {
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
	CAFile   string `yaml:"CAFile"`
}

type CCloud struct {
	Environment string `yaml:"environment" validate:"required"`
}

// Export
type Export struct {
	Topics struct {
		Exclude string `yaml:"exclude"`
	} `yaml:"topics"`
	Exporters []Exporter `yaml:"exporters"`
	Output    string     `yaml:"output" validate:"required"`
}

type Exporter string

const (
	Excel Exporter = "excel"
	Yaml  Exporter = "yaml"
	Json  Exporter = "json"
)
