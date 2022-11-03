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
	Resources []Resource `yaml:"resources" validate:"required"`
	Topics    Topics     `yaml:"topics" validate:"omitempty"`
	Exporters []Exporter `yaml:"exporters"`
	Output    string     `yaml:"output" validate:"required"`
}

type Topics struct {
	Exclude string `yaml:"exclude"`
	Include string `yaml:"include"`
}

type Resource string

const (
	ExportTopics         Resource = "topics"
	ExportSchemas        Resource = "schemas"
	ExportConsumerGroups Resource = "consumer_groups"
)

func (e Resource) String() string {
	resources := [...]string{"topics", "schemas", "consumer_groups"}

	x := string(e)
	for _, v := range resources {
		if v == x {
			return x
		}
	}

	return ""
}

type Exporter string

const (
	Excel Exporter = "excel"
	Yaml  Exporter = "yaml"
	Json  Exporter = "json"
)
