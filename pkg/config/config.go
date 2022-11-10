package config

type Config struct {
	Cluster         string         `yaml:"cluster" validate:"required"`
	EndpointUrl     string         `yaml:"endpointUrl" validate:"required"`
	BootstrapServer string         `yaml:"bootstrapServer" validate:"required"`
	Credentials     Credentials    `yaml:"credentials"`
	CCloud          CCloud         `yaml:"ccloud" validate:"omitempty"`
	Export          Export         `yaml:"export"`
	SchemaRegistry  SchemaRegistry `yaml:"schemaregistry" validate:"omitempty"`
}

type SchemaRegistry struct {
	EndpointUrl string      `yaml:"endpointUrl"`
	Credentials Credentials `yaml:"credentials"`
}

type Credentials struct {
	Key          string       `yaml:"key"`
	Secret       string       `yaml:"secret"`
	Certificates Certificates `yaml:"certificates" validate:"omitempty"`
}

type Certificates struct {
	CertFile string `yaml:"certFile"`
	KeyFile  string `yaml:"keyFile"`
	CAFile   string `yaml:"CAFile"`
}

type CCloud struct {
	Environment string `yaml:"environment" validate:"omitempty"`
}

// EXPORT Configuration
// **********************
type Export struct {
	Resources []Resource `yaml:"resources" validate:"required"`
	Topics    Topics     `yaml:"topics" validate:"omitempty"`
	CLink     CLink      `yaml:"clink" validate:"omitempty"`
	CFK       CFK        `yaml:"cfk" validate:"omitempty"`
	Exporters []Exporter `yaml:"exporters"`
	Output    string     `yaml:"output" validate:"required"`
	Schemas   Schemas    `yaml:"schemas" validate:"omitempty"`
}

// Topics export configuration
type Topics struct {
	Exclude string `yaml:"exclude"`
	Include string `yaml:"include"`
}

// Clink export configuration
type CLink struct {
	Name        string `yaml:"name"`
	Destination string `yaml:"destination"`
	AutoCreate  bool   `yaml:"autocreate"`
	Sync        struct {
		Offset bool `yaml:"offset"`
		Acl    bool `yaml:"acl"`
	} `yaml:"sync"`
}

// CFK export configuration
type CFK struct {
	Namespace      string `yaml:"namespace"`
	KafkaRestClass string `yaml:"kafkarestclass"`
}

// Schema registry
type Schemas struct {
	Subject struct {
		Version SubjectVersion `yaml:"version"`
	} `yaml:"subject"`
}

type SubjectVersion string

const (
	Latest SubjectVersion = "latest"
	All    SubjectVersion = "all"
)

// Resources
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
	Clink Exporter = "clink"
	Cfk   Exporter = "cfk"
)
