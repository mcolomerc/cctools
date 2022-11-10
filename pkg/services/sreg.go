package services

import (
	"encoding/json"
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/model"
	"mcolomerc/cc-tools/pkg/export/sregexp"
)

type SchemasService struct {
	RestClient        client.RestClient
	Conf              config.Config
	SchemaRegistryUrl string
	Exporters         []sregexp.SRegExporter
}

func NewSchemasService(conf config.Config) *SchemasService {
	restClient := client.New(conf.SchemaRegistry.EndpointUrl, conf.SchemaRegistry.Credentials)
	var exporters []sregexp.SRegExporter
	for _, v := range conf.Export.Exporters {
		if v == config.Excel {
			exporters = append(exporters, &sregexp.SRegExcelExporter{})
		} else if v == config.Json {
			exporters = append(exporters, &sregexp.SRegJsonExporter{})
		} else if v == config.Yaml {
			exporters = append(exporters, &sregexp.SRegYamlExporter{})
		} else {
			log.Printf("Schema Registry exporter: Unrecognized exporter: %v \n", v)
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
	outputPath := service.Conf.Export.Output + "/" + service.Conf.Cluster

	for _, v := range service.Conf.Export.Resources {
		if v == config.ExportSchemas {
			log.Print("Exporting Subjects Info")
			result := service.GetSubjects()
			done := make(chan bool, len(exportExecutors))
			for _, v := range exportExecutors {
				go func(v sregexp.SRegExporter) {
					err := v.ExportSubjects(result, outputPath)
					if err != nil {
						log.Printf("Error: %s\n", err)
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
		log.Printf("Error getting Schema Registry config : %s\n", err)
	}
	return config
}

func (service *SchemasService) GetSubjects() []model.SubjectVersion {
	subjects, err := service.RestClient.GetList(service.SchemaRegistryUrl + "subjects")
	if err != nil {
		log.Printf("Error getting Schema Registry config : %s\n", err)
	}
	resp := make([]model.SubjectVersion, len(subjects))
	done := make(chan []model.SubjectVersion, len(subjects))
	for _, v := range subjects {
		go func(subj string) {
			done <- service.GetSubjectVersions(subj)
		}(fmt.Sprint(v))
	}
	for i := 0; i < len(subjects); i++ {
		r := <-done
		resp = append(resp, r...)
	}
	return resp
}

func (service *SchemasService) GetSubjectVersions(subject string) []model.SubjectVersion {
	subjectsVersions, err := service.RestClient.GetList(service.SchemaRegistryUrl + "subjects/" + subject + "/versions")
	if err != nil {
		log.Printf("Error getting Schema Registry GetSubjectVersions : %s\n", err)
	}
	resp := make([]model.SubjectVersion, len(subjectsVersions))
	done := make(chan model.SubjectVersion, len(subjectsVersions))
	for _, v := range subjectsVersions {
		id := int(v.(float64))
		go func(subject string, version string) {
			done <- service.GetSubjectVersion(subject, version)
		}(subject, fmt.Sprintf("%d", id))
	}
	for i := 0; i < len(subjectsVersions); i++ {
		resp = append(resp, <-done)
	}
	close(done)
	return resp
}

func (service *SchemasService) GetSubjectVersion(subject string, version string) model.SubjectVersion {
	subjVersion, err := service.RestClient.Get(service.SchemaRegistryUrl + "subjects/" + subject + "/versions/" + version)
	if err != nil {
		log.Printf("Error getting Schema Registry GetSubjectVersion : %s\n", err)
	}
	data := subjVersion.(map[string]interface{})
	jsonString, _ := json.Marshal(data)
	subjectVersion := &model.SubjectVersion{}
	json.Unmarshal([]byte(jsonString), &subjectVersion)
    
	return *subjectVersion
}

func (service *SchemasService) GetSubjectConfig(subject string) interface{} {
	subjectConfig, err := service.RestClient.Get(service.SchemaRegistryUrl + "subjects/" + subject + "/config")
	if err != nil {
		log.Printf("Error getting Schema Registry config : %s\n", err)
	}
	return subjectConfig
}

func (service *SchemasService) GetSchemas() interface{} {
	subjects, err := service.RestClient.Get(service.SchemaRegistryUrl + "schemas")
	if err != nil {
		log.Printf("Error getting Schema Registry GetSchemas : %s\n", err)
	}
	return subjects
}
