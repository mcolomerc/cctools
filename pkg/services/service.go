package services

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
)

type Service interface {
	Export()
}

type ExportHandler struct {
	Services map[string]Service
}

const (
	KAFKA_SERVICE   = "kafka"
	SCHEMAS_SERVICE = "schemas"
	GIT_SERVICE     = "git"
)

func NewExportHandler(conf config.Config) *ExportHandler {
	services := make(map[string]Service)
	for _, resource := range conf.Export.Resources {
		if resource == config.ExportTopics {
			services[KAFKA_SERVICE] = NewKafkaService(conf)
		}
		if resource == config.ExportConsumerGroups {
			if _, ok := services[KAFKA_SERVICE]; !ok {
				services[KAFKA_SERVICE] = NewKafkaService(conf)
			}
		}
		if resource == config.ExportSchemas {
			services[SCHEMAS_SERVICE] = NewSchemasService(conf)
		}
	}
	log.Debug("Check Git repositories...", len(conf.Export.Git))
	log.Debug("Check Git repositories...", conf.Export.Git)
	if len(conf.Export.Git) > 0 {
		services[GIT_SERVICE] = NewGitService(conf)
	}
	return &ExportHandler{
		Services: services,
	}
}

func (exp *ExportHandler) BuildExport() {
	done := make(chan bool, len(exp.Services))
	for _, v := range exp.Services {
		go func(s Service) {
			s.Export()
			done <- true
		}(v)
	}
	for i := 0; i < len(exp.Services); i++ {
		<-done
	}
	close(done)
}
