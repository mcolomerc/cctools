package ccloud

type Topic struct {
	Name              string
	Partitions        interface{}
	ReplicationFactor interface{}
	Configs           []TopicConfig
}

type TopicConfig struct {
	Name  string
	Value interface{}
}
