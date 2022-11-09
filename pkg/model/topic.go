package model

type Topic struct {
	Name              string         `yaml:"name" json:"name"`
	Partitions        interface{}    `yaml:"partitions" json:"partitions"`
	ReplicationFactor interface{}    `yaml:"replicationFactor" json:"replicationFactor"`
	MinIsr            interface{}    `yaml:"minIsr" json:"minIsr"`
	RetentionTime     interface{}    `yaml:"retentionTime" json:"retentionTime"`
	Configs           []TopicConfig  `yaml:"configs" json:"configs"`
	RoleBindings      []RoleBinding  `yaml:"roleBindings" json:"roleBindings"`
}

type TopicConfig struct {
	Name  string      `yaml:"name" json:"name"`
	Value interface{} `yaml:"value" json:"value"`
}
type RoleBinding struct {
	RoleName string
	Users interface{}
}