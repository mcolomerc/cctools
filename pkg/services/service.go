package services

import (
	"mcolomerc/cc-tools/pkg/config"
)

type Service interface {
	Export()
}

type ExportHandler struct {
	Services []Service
}

func NewExportHandler(conf config.Config) *ExportHandler {
	var services []Service
	for _, resource := range conf.Export.Resources {
		if resource == config.ExportTopics || resource == config.ExportConsumerGroups {
			services = append(services, NewKafkaService(conf))
		}
		if resource == config.ExportSchemas {
			services = append(services, NewSchemasService(conf))
		}
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
