package model

type Topic struct {
	Name              string        `yaml:"name" json:"name"`
	Partitions        interface{}   `yaml:"partitions" json:"partitions"`
	ReplicationFactor interface{}   `yaml:"replicationFactor" json:"replicationFactor"`
	Configs           []TopicConfig `yaml:"configs" json:"configs"`
}

type TopicConfig struct {
	Name  string      `yaml:"name" json:"name"`
	Value interface{} `yaml:"value" json:"value"`
}
