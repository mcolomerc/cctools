package imprt

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
)

type ImportService interface {
	Import()
}

type ImportHandler struct {
	Services map[string]ImportService
}

const (
	KAFKA_SERVICE   = "kafka"
	SCHEMAS_SERVICE = "schemas"
	GIT_SERVICE     = "git"
)

func NewImportHandler(conf config.Config) (*ImportHandler, error) {
	services := make(map[string]ImportService)
	for _, resource := range conf.Import.Resources {
		if resource == config.Topic || resource == config.ConsumerGroup {
			serv, err := NewKafkaImport(conf)
			if err != nil {
				log.Error("Error creating Kafka service: ", err)
				return nil, err
			}
			services[KAFKA_SERVICE] = serv
		}
		if resource == config.Schema {
			srserv, err := NewSchemaImport(conf)
			if err != nil {
				log.Error("Error creating Schema Registry service: " + err.Error())
				return nil, err
			}
			services[SCHEMAS_SERVICE] = srserv
		}
	}

	return &ImportHandler{
		Services: services,
	}, nil
}

func (exp *ImportHandler) Import() {
	done := make(chan bool, len(exp.Services))
	for _, v := range exp.Services {
		go func(s ImportService) {
			s.Import()
			done <- true
		}(v)
	}
	for i := 0; i < len(exp.Services); i++ {
		<-done
	}
	close(done)
}
