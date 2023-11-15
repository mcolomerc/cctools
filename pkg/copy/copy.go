package copy

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
)

type CopyService interface {
	Copy()
}

type CopyHandler struct {
	Services map[string]CopyService
}

const (
	COPY_KAFKA_SERVICE   = "kafka"
	COPY_SCHEMAS_SERVICE = "schemaregistry"
)

func NewCopyHandler(conf config.Config) (*CopyHandler, error) {
	services := make(map[string]CopyService)
	for _, resource := range conf.Export.Resources {
		if resource == config.Topic || resource == config.ConsumerGroup {
			serv, err := NewKafkaCopy(conf)
			if err != nil {
				log.Error("Error creating Kafka copy service: ", err)
				return nil, err
			}
			services[COPY_KAFKA_SERVICE] = serv
		}
		if resource == config.Schema {
			srserv, err := NewSchemaCopy(conf)
			if err != nil {
				log.Error("Error creating Schema Registry service: " + err.Error())
				return nil, err
			}
			services[COPY_SCHEMAS_SERVICE] = srserv
		}
	}
	return &CopyHandler{
		Services: services,
	}, nil
}

func (exp *CopyHandler) Copy() {
	done := make(chan bool, len(exp.Services))
	for _, v := range exp.Services {
		go func(s CopyService) {
			s.Copy()
			done <- true
		}(v)
	}
	for i := 0; i < len(exp.Services); i++ {
		<-done
	}
	close(done)
}
