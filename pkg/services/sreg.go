package services

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/export"
)

type SchemasService struct {
	RestClient        client.RestClient
	Conf              config.Config
	SchemaRegistryUrl string
	Exporters         []export.Exporter
}

func NewSchemasService(conf config.Config) *SchemasService {
	restClient := client.New(conf.SchemaRegistry.EndpointUrl, conf.SchemaRegistry.Credentials)
	var exporters []export.Exporter
	for _, v := range conf.Export.Exporters {
		if v == config.Json {
			exporters = append(exporters, &export.JsonExporter{})
		} else if v == config.Yaml {
			exporters = append(exporters, &export.YamlExporter{})
		} else {
			fmt.Printf("Schema Registry exporter: Unrecognized exporter: %v \n", v)
		}
	}
	return &SchemasService{
		RestClient:        *restClient,
		Conf:              conf,
		SchemaRegistryUrl: fmt.Sprintf("%s/", conf.SchemaRegistry.EndpointUrl),
		Exporters:         exporters,
	}
}

func (service *SchemasService) Export() {
	exportExecutors := service.Exporters
	outputPath := service.Conf.Export.Output + "/_subjects"

	for _, v := range service.Conf.Export.Resources {
		if v == config.ExportSchemas {
			result := service.GetSubjects()
			done := make(chan bool, len(exportExecutors))
			for _, v := range exportExecutors {
				go func(v export.Exporter) {
					err := v.Export(result, outputPath)
					if err != nil {
						fmt.Printf("Error: %s\n", err)
					}
					done <- true
				}(v)
			}
			for i := 0; i < len(exportExecutors); i++ {
				<-done
			}
			close(done)
		}
	}
}

func (service *SchemasService) GetConfig() interface{} {
	config, err := service.RestClient.Get(service.SchemaRegistryUrl + "config")
	if err != nil {
		log.Printf("client: error getting Schema Registry config : %s\n", err)
	}
	return config
}

func (service *SchemasService) GetSubjects() interface{} {
	subjects, err := service.RestClient.Get(service.SchemaRegistryUrl + "subjects")
	if err != nil {
		log.Printf("client: error getting Schema Registry config : %s\n", err)
	}
	return subjects
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
