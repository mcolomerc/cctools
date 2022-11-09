package model

type Scope struct {
	Clusters Clusters `json:"clusters"`
}

type Clusters struct {
	Kafka          string `json:"kafka-cluster"`
	SchemaRegistry string `json:"schemaregistry-cluster" validate:"omitempty"`
	Ksql           string `json:"ksql-cluster" validate:"omitempty"`
}