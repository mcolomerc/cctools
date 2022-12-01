package kafkaexp

import (
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"

	"gopkg.in/yaml.v2"
)

type KafkaCfkExporter struct {
	ParentKafkaExporter
	export.CfkExporter
}

type KafkaRestClass struct {
	Name string `yaml:"name"`
}
type Spec struct {
	Replicas       interface{}            `yaml:"replicas"`
	Partitions     interface{}            `yaml:"partitionCount"`
	Configs        map[string]interface{} `yaml:"configs"`
	KafkaRestClass KafkaRestClass         `yaml:"kafkaRestClassRef"`
}
type Topic_CRD struct {
	model.CRD `yaml:",inline"`
	Spec      Spec `yaml:"spec"`
}

const (
	apiVersion = "platform.confluent.io/v1beta1"
	kind       = "KafkaTopic"
)

func NewKafkaCfkExporter(config config.Config) *KafkaCfkExporter {
	return &KafkaCfkExporter{
		CfkExporter: export.CfkExporter{
			Namespace:  config.Export.CFK.Namespace,
			ClusterRef: config.Export.CFK.KafkaRestClass,
		},
	}

}

func (e KafkaCfkExporter) ExportTopics(topics []model.Topic, outputPath string) error {

	for _, v := range topics {
		configs := make(map[string]interface{})
		for _, vc := range v.Configs {
			configs[vc.Name] = vc.Value
		}
		c := &model.CRD{
			ApiVersion: apiVersion,
			Kind:       kind,
			Metadata: export.Metadata{
				Name:      v.Name,
				Namespace: e.Namespace,
			},
		}
		crd := &Topic_CRD{
			CRD: *c,
			Spec: Spec{
				Replicas:   v.ReplicationFactor,
				Partitions: v.Partitions,
				Configs:    configs,
				KafkaRestClass: KafkaRestClass{
					Name: e.ClusterRef,
				},
			},
		}
		file, errJson := yaml.Marshal(crd)
		if errJson != nil {
			return errJson
		}
		err := ioutil.WriteFile(outputPath+"/"+v.Name+".yml", file, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
