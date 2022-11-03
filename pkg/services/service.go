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

func NewExportHandler(conf config.RuntimeConfig) *ExportHandler {
	var services []Service
	services = append(services, NewKafkaService(conf))
	return &ExportHandler{
		Services: services,
	}
}

func (exp *ExportHandler) BuildExport() {
	for _, v := range exp.Services {
		v.Export()
	}
}
