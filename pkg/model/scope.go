package model

type Scope struct {
	Clusters Clusters `json:"clusters" validate:"required"`
}

type Clusters struct {
	Kafka          string `json:"kafka-cluster" validate:"required"`
	SchemaRegistry string `json:"schemaregistry-cluster,omitempty" validate:"omitempty"`
	Ksql           string `json:"ksql-cluster,omitempty" validate:"omitempty"`
}