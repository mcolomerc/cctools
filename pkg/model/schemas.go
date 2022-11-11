package model

type Schema struct {
	Subject    string `yaml:"subject" json:"subject"`
	Version    int    `yaml:"version" json:"version"`
	Id         int    `yaml:"id" json:"id"`
	SchemaType string `yaml:"schemaType" json:"schemaType"`
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
