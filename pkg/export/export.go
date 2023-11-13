package export

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
)

type ExportService interface {
	Export()
}

type ExportHandler struct {
	Services map[string]ExportService
}

const (
	KAFKA_SERVICE   = "kafka"
	SCHEMAS_SERVICE = "schemas"
	GIT_SERVICE     = "git"
)

func NewExportHandler(conf config.Config) (*ExportHandler, error) {
	services := make(map[string]ExportService)
	for _, resource := range conf.Export.Resources {
		if resource == config.Topic || resource == config.ConsumerGroup {
			serv, err := NewKafkaService(conf)
			if err != nil {
				log.Error("Error creating Kafka service: ", err)
				return nil, err
			}
			services[KAFKA_SERVICE] = serv
		}
		if resource == config.Schema {
			srserv, err := NewSchemasService(conf)
			if err != nil {
				log.Error("Error creating Schema Registry service: " + err.Error())
				return nil, err
			}
			services[SCHEMAS_SERVICE] = srserv
		}
	}

	log.Debug("Check Git repositories...", conf.Export.Git)
	if len(conf.Export.Git) > 0 {
		services[GIT_SERVICE] = NewGitService(conf)
	}
	return &ExportHandler{
		Services: services,
	}, nil
}

func (exp *ExportHandler) BuildExport() {
	done := make(chan bool, len(exp.Services))
	for _, v := range exp.Services {
		go func(s ExportService) {
			s.Export()
			done <- true
		}(v)
	}
	for i := 0; i < len(exp.Services); i++ {
		<-done
	}
	close(done)
}
