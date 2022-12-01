package sregexp

import (
	"fmt"
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
	"strings"

	"gopkg.in/yaml.v2"
)

type SRegCfkExporter struct {
	export.CfkExporter
}

type SchemaRegistryClusterRef struct {
	Name string `yaml:"name"`
}
type Data struct {
	ConfigReg string `yaml:"configRef"`
	Format    string `yaml:"format"`
}
type Spec struct {
	Data                     `yaml:"data"`
	SchemaRegistryClusterRef `yaml:"schemaRegistryClusterRef"`
}

type Schema struct {
	model.CRD `yaml:",inline"`
	Spec      Spec `yaml:"spec"`
}

type ConfigMap struct {
	model.CRD  `yaml:",inline"`
	ConfigData `yaml:"data"`
}

type ConfigData struct {
	Schema string `yaml:"schema"`
}

const (
	apiVersion       = "platform.confluent.io/v1beta1"
	kind             = "Schema"
	configKind       = "ConfigMap"
	configApiVersion = "v1"
	configSchema     = " | \n"
)

func NewSRegCfkExporter(config config.Config) *SRegCfkExporter {
	return &SRegCfkExporter{
		CfkExporter: export.CfkExporter{
			Namespace:  config.Export.CFK.Namespace,
			ClusterRef: config.Export.CFK.SchemaRegistryClusterRef,
		},
	}
}

func (e SRegCfkExporter) ExportSchemas(schemas []model.Schema, outputPath string) error {
	done := make(chan bool, len(schemas))
	for _, s := range schemas {
		go func(schema model.Schema, outputPath string) {
			configOut := fmt.Sprintf("%s%s_%d", outputPath, schema.Subject+"-config", schema.Version)
			configCrd := &model.CRD{
				ApiVersion: configApiVersion,
				Kind:       configKind,
				Metadata: export.Metadata{
					Name:      schema.Subject + "-config",
					Namespace: e.Namespace,
				},
			}
			configData := &ConfigData{
				Schema: schema.Schema,
			}

			configMap := &ConfigMap{
				CRD:        *configCrd,
				ConfigData: *configData,
			}

			fileConfig, _ := yaml.Marshal(configMap)
			ioutil.WriteFile(configOut+".yml", fileConfig, 0644)

			out := fmt.Sprintf("%s%s_%d", outputPath, schema.Subject, schema.Version)
			schemaCrd := &model.CRD{
				ApiVersion: apiVersion,
				Kind:       kind,
				Metadata: export.Metadata{
					Name:      schema.Subject,
					Namespace: e.Namespace,
				},
			}
			schemaType := "avro"
			if len(schema.SchemaType) > 0 {
				schemaType = strings.ToLower(schema.SchemaType)
			}
			crd := &Schema{
				CRD: *schemaCrd,
				Spec: Spec{
					Data: Data{
						ConfigReg: schema.Subject + "-config",
						Format:    schemaType,
					},
					SchemaRegistryClusterRef: SchemaRegistryClusterRef{
						Name: e.ClusterRef,
					},
				},
			}
			file, _ := yaml.Marshal(crd)
			ioutil.WriteFile(out+".yml", file, 0644)
			done <- true
		}(s, outputPath)
	}
	for i := 0; i < len(schemas); i++ {
		<-done
	}
	close(done)
	return nil
}

func (e SRegCfkExporter) ExportSubjects(schema []model.SubjectVersion, outputPath string) error {
	return nil
}
