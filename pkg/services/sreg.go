package services

import (
	"fmt"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
)

type SchemasService struct {
	RestClient        client.RestClient
	Conf              config.Config
	SchemaRegistryUrl string
}

func NewSchemasService(conf config.Config) *SchemasService {
	restClient := client.New(conf)
	return &SchemasService{
		RestClient:        *restClient,
		Conf:              conf,
		SchemaRegistryUrl: fmt.Sprintf("%s/", conf.SchemaRegistryUrl),
	}
}

func (service *SchemasService) Export() {
}

func (service *SchemasService) GetSubjects() {
	// /subjects
}

func (service *SchemasService) GetSubjectVersions() {
	// /subjects
}

/**
List all schema versions registered under the subject “Kafka-value”
curl -X GET http://localhost:8081/subjects/Kafka-value/versions
Example result:

[1]
Fetch Version 1 of the schema registered under subject “Kafka-value”
curl -X GET http://localhost:8081/subjects/Kafka-value/versions/1
Example result:

{"subject":"Kafka-value","version":1,"id":1,"schema":"\"string\""}

Fetch the schema again by globally unique ID 1
curl -X GET http://localhost:8081/schemas/ids/1

Get the top level config
curl -X GET http://localhost:8081/config

Get compatibility requirements on a subject
curl -X GET http://localhost:8081/config/my-kafka-value

List schema types currently registered in Schema Registry
curl -X GET http://localhost:8081/schemas/types

List all subject-version pairs where a given ID is used
curl -X GET http://localhost:8081/schemas/ids/2/versions

List IDs of schemas that reference a given schema
curl -X GET http://localhost:8081/subjects/other.proto/versions/1/referencedby


**/
