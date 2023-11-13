package export

import (
	"encoding/json"
	"fmt"
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/exporters"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
	"mcolomerc/cc-tools/pkg/util"
)

type SchemasService struct {
	RestClient        client.RestClient
	Conf              config.Config
	SchemaRegistryUrl string
	SRegExporters     []exporters.SRegExporter
	Paths             SRPaths
}

type SRPaths struct {
	Schemas  string
	Subjects string
}

const (
	SCHEMAS_PATH = "/schemas/"
	SUBJECT_PATH = "/subjects/"
)

func NewSchemasService(conf config.Config) (*SchemasService, error) {
	restClient := client.NewRestClient(conf.Source.SchemaRegistry.EndpointUrl, conf.Source.SchemaRegistry.Credentials)
	var exportersList []exporters.SRegExporter
	for _, v := range conf.Export.Exporters {
		if v == config.Json {
			exportersList = append(exportersList, exporters.NewSRegJsonExporter())
		} else if v == config.Yaml {
			exportersList = append(exportersList, exporters.NewSRegYamlExporter())
		} else if v == config.Excel {
			exportersList = append(exportersList, exporters.NewSRegExcelExporter())
		} else if v == config.Cfk {
			exportersList = append(exportersList, exporters.NewSRegCfkExporter(conf))
		} else {
			log.Error("Schema Registry exporter: Unrecognized exporter: " + v)
			return nil, fmt.Errorf("Schema Registry exporter - Unrecognized exporter: %s", v)
		}
	}
	paths := &SRPaths{
		Schemas:  conf.Export.Output + SCHEMAS_PATH,
		Subjects: conf.Export.Output + SUBJECT_PATH,
	}
	return &SchemasService{
		RestClient:        *restClient,
		Conf:              conf,
		SchemaRegistryUrl: fmt.Sprintf("%s/", conf.Source.SchemaRegistry.EndpointUrl),
		SRegExporters:     exportersList,
		Paths:             *paths,
	}, nil
}

func (service *SchemasService) buildExportPaths() {
	util.BuildPath(service.Paths.Schemas)
	util.BuildPath(service.Paths.Subjects)
}

func (service *SchemasService) Export() {
	log.Debug("Exporting from Schema Registry...")
	service.buildExportPaths()
	done := make(chan bool, 2)
	for _, v := range service.Conf.Export.Resources {
		if v == config.Schema {
			go service.exportSubjects(done)
			go service.exportSchemas(done)
		}
	}
	for i := 0; i < 2; i++ {
		<-done
	}
	close(done)
	log.Info("Schema registry exported: Schemas: " + service.Paths.Schemas + " Subjects: " + service.Paths.Subjects)
}

func (service *SchemasService) exportSchemas(exported chan bool) {
	exportExecutors := service.SRegExporters
	result := service.GetSchemas()
	done := make(chan bool, len(exportExecutors))
	for _, v := range exportExecutors {
		go func(v exporters.SRegExporter, s []model.Schema) {
			pth := fmt.Sprintf("%s%s/", service.Paths.Schemas, v.GetPath())
			log.Debug("Building path :: ", pth)
			util.BuildPath(pth)
			err := v.ExportSchemas(s, pth)
			if err != nil {
				log.Error("Error: %s\n", err)
			}
			done <- true
		}(v, result)
	}
	for i := 0; i < len(exportExecutors); i++ {
		<-done
	}
	close(done)
	exported <- true
}
func (service *SchemasService) exportSubjects(exported chan bool) {
	exportExecutors := service.SRegExporters
	result := service.GetSubjects()
	done := make(chan bool, len(exportExecutors))
	for _, v := range exportExecutors {
		go func(v exporters.SRegExporter, s []model.SubjectVersion) {
			pth := fmt.Sprintf("%s%s/", service.Paths.Subjects, v.GetPath())
			log.Debug("Building path :: ", pth)
			util.BuildPath(pth)
			err := v.ExportSubjects(s, pth)
			if err != nil {
				log.Error("Error: %s\n", err)
			}
			done <- true
		}(v, result)
	}
	for i := 0; i < len(exportExecutors); i++ {
		<-done
	}
	close(done)
	exported <- true
}

func (service *SchemasService) GetConfig() interface{} {
	config, err := service.RestClient.Get(service.SchemaRegistryUrl + "config")
	if err != nil {
		log.Error("Error getting Schema Registry config : " + err.Error())
	}
	return config
}

func (service *SchemasService) GetSubjects() []model.SubjectVersion {
	subjects, err := service.RestClient.GetList(service.SchemaRegistryUrl + "subjects")
	if err != nil {
		log.Error("Error getting Schema Registry config : ", err.Error())
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
func (service *SchemasService) Copy() {
	log.Debug("Copying Schema Registry...")
	subjects := service.GetSubjects()
	done := make(chan bool, len(subjects))
	for _, subject := range subjects {
		go func(subject model.SubjectVersion) {
			jsonsubject, err := json.Marshal(subject)
			if err != nil {
				log.Error("Error marshalling subject versions : " + err.Error())
			}
			_, err = service.RestClient.Post(service.SchemaRegistryUrl+"subjects", jsonsubject)
			if err != nil {
				log.Error("Error post Schema Registry config : " + err.Error())
			}
			done <- true
		}(subject)
	}
	for i := 0; i < len(subjects); i++ {
		<-done
	}
	close(done)
}
func (service *SchemasService) GetSubjectVersions(subject string) []model.SubjectVersion {
	subjectsVersions, err := service.RestClient.GetList(service.SchemaRegistryUrl + "subjects/" + subject + "/versions")
	if err != nil {
		log.Error("Error getting Schema Registry GetSubjectVersions : " + subject)
		log.Error(err)
	}
	resp := make([]model.SubjectVersion, len(subjectsVersions))
	done := make(chan model.SubjectVersion, len(subjectsVersions))

	// Generate []int array
	versions := make([]int, len(subjectsVersions))
	for i := range subjectsVersions {
		versions[i] = int(subjectsVersions[i].(float64))
	}

	var latest int
	if service.Conf.Export.Schemas.Subject.Version == config.Latest {
		latest = getLatestVersion(versions)
		versions = []int{latest}
	}

	for _, v := range versions {
		go func(subject string, version string) {
			done <- service.GetSubjectVersion(subject, version)
		}(subject, fmt.Sprintf("%d", v))
	}
	for i := 0; i < len(versions); i++ {
		resp = append(resp, <-done)
	}
	close(done)
	return resp
}

func (service *SchemasService) GetSubjectVersion(subject string, version string) model.SubjectVersion {
	subjVersion, err := service.RestClient.Get(service.SchemaRegistryUrl + "subjects/" + subject + "/versions/" + version)
	if err != nil {
		log.Error("Error getting Schema Registry GetSubjectVersion : " + err.Error())
	}
	data := subjVersion.(map[string]interface{})
	jsonString, _ := json.Marshal(data)
	subjectVersion := &model.SubjectVersion{}
	json.Unmarshal([]byte(jsonString), &subjectVersion)
	return *subjectVersion
}

func (service *SchemasService) GetSubjectConfig(subject string) model.CompatibilityMode {
	subjectConfig, err := service.RestClient.Get(service.SchemaRegistryUrl + "config/" + subject)
	if err != nil {
		log.Error("Error getting Schema Registry config : " + err.Error())
	}
	data := subjectConfig.(map[string]interface{})
	jsonString, _ := json.Marshal(data)
	compatibilityMode := &model.CompatibilityMode{}
	json.Unmarshal([]byte(jsonString), &compatibilityMode)
	return *compatibilityMode
}

func (service *SchemasService) GetSchemas() []model.Schema {
	schemas, err := service.RestClient.GetList(service.SchemaRegistryUrl + "schemas")
	if err != nil {
		log.Error("Error getting Schema Registry GetSchemas : %s\n", err)
	}

	var resp []model.Schema
	done := make(chan model.Schema, len(schemas))
	for _, v := range schemas {
		go func(v interface{}) {
			data := v.(map[string]interface{})
			jsonString, _ := json.Marshal(data)
			schema := &model.Schema{}
			json.Unmarshal([]byte(jsonString), &schema)
			done <- *schema
		}(v)
	}
	for i := 0; i < len(schemas); i++ {
		resp = append(resp, <-done)
	}
	close(done)
	if service.Conf.Export.Schemas.Version == config.Latest {
		elementMap := make(map[string]model.Schema)
		for _, v := range resp {
			if val, ok := elementMap[v.Subject]; ok {
				if v.Version > val.Version {
					elementMap[v.Subject] = v
				}
			} else {
				elementMap[v.Subject] = v
			}
		}
		var elementsArray []model.Schema
		for _, v := range elementMap {
			elementsArray = append(elementsArray, v)
		}
		return elementsArray
	}
	return resp
}

func getLatestVersion(versions []int) int {
	var max int
	for i, e := range versions {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}
