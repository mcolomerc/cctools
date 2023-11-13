package imprt

import (
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
)

type SchemaImport struct {
	DestClient client.RestClient
	Conf       config.Config
	Paths      SchemaPaths
}

type SchemaPaths struct {
	Schemas  string
	Subjects string
}

const (
	SCHEMAS_PATH  = "/schemas/"
	SUBJECTS_PATH = "/subjects/"
)

func NewSchemaImport(conf config.Config) (*SchemaImport, error) {
	paths := &SchemaPaths{
		Schemas:  conf.Export.Output + TOPICS_PATH,
		Subjects: conf.Export.Output + CGROUPS_PATH,
	}
	destClient := client.NewRestClient(conf.Source.SchemaRegistry.EndpointUrl, conf.Source.SchemaRegistry.Credentials)

	kafkaService := &SchemaImport{
		DestClient: *destClient,
		Conf:       conf,
		Paths:      *paths,
	}

	return kafkaService, nil
}

func (k *SchemaImport) Import() {
}
