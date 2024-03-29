package model

type Topic struct {
	Name              string        `yaml:"name" json:"name"`
	Partitions        int           `yaml:"partitions" json:"partitions"`
	ReplicationFactor int           `yaml:"replicationFactor" json:"replicationFactor"`
	MinIsr            string        `yaml:"minIsr" json:"minIsr"`
	RetentionTime     string        `yaml:"retentionTime" json:"retentionTime"`
	Configs           []TopicConfig `yaml:"configs" json:"configs"`
	RoleBindings      []RoleBinding `yaml:"roleBindings" json:"roleBindings"`
	ACLs              []AclBinding  `yaml:"acls" json:"acls"`
}

type AclBinding struct {
	Principal    string `yaml:"principal" json:"principal"`
	Host         string `yaml:"host" json:"host"`
	Operation    string `yaml:"operation" json:"operation"`
	Permission   string `yaml:"permission" json:"permission"`
	ResourceType string `yaml:"resourceType" json:"resourceType"`
	ResourceName string `yaml:"resourceName" json:"resourceName"`
	PatternType  string `yaml:"patternType" json:"patternType"`
}

type TopicConfig struct {
	Name  string      `yaml:"name" json:"name"`
	Value interface{} `yaml:"value" json:"value"`
}
type RoleBinding struct {
	RoleName string
	Users    interface{}
}
