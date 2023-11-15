package model

type Schema struct {
	Subject    string `yaml:"subject" json:"subject,omitempty"`
	Version    int    `yaml:"version" json:"version,omitempty"`
	Id         int    `yaml:"id" json:"id,omitempty"`
	SchemaType string `yaml:"schemaType" json:"schemaType,omitempty"`
	Schema     string `yaml:"schema" json:"schema"`
}

type SubjectVersion struct {
	Subject string `yaml:"subject" json:"subject"`
	Version int    `yaml:"version" json:"version"`
	Id      int    `yaml:"id" json:"id"`
	Schema  string `yaml:"schema" json:"schema"`
}

type CompatibilityMode struct {
	Mode string `yaml:"compatibilityLevel" json:"compatibilityLevel"`
}

type SchemaRegistryMode string

const (
	Import    SchemaRegistryMode = "IMPORT"
	ReadWrite SchemaRegistryMode = "READWRITE"
)

func (e SchemaRegistryMode) String() string {
	resources := [...]string{"IMPORT", "READWRITE"}

	x := string(e)
	for _, v := range resources {
		if v == x {
			return x
		}
	}

	return ""
}
