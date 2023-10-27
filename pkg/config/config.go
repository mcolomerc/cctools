package config

type Config struct {
	Cluster string `yaml:"cluster"`
	// KafkaAdmin client
	Source `yaml:"source" validate:"required"`
	// Getting rBAC from CC
	CCloud `yaml:"ccloud" validate:"omitempty"`
	// Export Configuration
	Export `yaml:"export"`
	// Schema Registry client
	SchemaRegistry `yaml:"schemaregistry" validate:"omitempty"`
}

type Source struct {
	BootstrapServer string            `yaml:"bootstrapServer"`
	ClientProps     map[string]string `yaml:"clientProps"`
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
	Git       map[string]string `yaml:"git"`
	Resources []Resource        `yaml:"resources"`
	Topics    `yaml:"topics" validate:"omitempty"`
	CLink     `yaml:"clink" validate:"omitempty"`
	CFK       `yaml:"cfk" validate:"omitempty"`
	Exporters []Exporter `yaml:"exporters"`
	Output    string     `yaml:"output" validate:"required"`
	Schemas   `yaml:"schemas" validate:"omitempty"`
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
	Namespace                string `yaml:"namespace"`
	KafkaRestClass           string `yaml:"kafkarestclass"`
	SchemaRegistryClusterRef string `yaml:"schemaRegistryClusterRef"`
}

// Schema registry
type Schemas struct {
	Version Version `yaml:"version"`
	Subject struct {
		Version Version `yaml:"version"`
	} `yaml:"subjects"`
}

type Version string

const (
	Latest Version = "latest"
	All    Version = "all"
)

// Resources
type Resource string

const (
	ExportTopics         Resource = "topics"
	ExportSchemas        Resource = "schemas"
	ExportConsumerGroups Resource = "consumer_groups"
	ExportAcls           Resource = "acls"
)

func (e Resource) String() string {
	resources := [...]string{"topics", "schemas", "consumer_groups", "acls"}

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
	Hcl   Exporter = "hcl"
)

func (e Exporter) String() string {
	resources := [...]string{"clink", "cfk", "hcl", "json", "yaml", "excel"}

	x := string(e)
	for _, v := range resources {
		if v == x {
			return x
		}
	}

	return ""
}
