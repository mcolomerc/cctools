package exporters

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/model"

	"os"

	"gopkg.in/yaml.v2"
)

type KafkaCfkExporter struct {
	ParentKafkaExporter
	CfkExporter
}

type KafkaRestClass struct {
	Name string `yaml:"name"`
}
type TopicSpec struct {
	Replicas       interface{}            `yaml:"replicas"`
	Partitions     interface{}            `yaml:"partitionCount"`
	Configs        map[string]interface{} `yaml:"configs"`
	KafkaRestClass KafkaRestClass         `yaml:"kafkaRestClassRef"`
}

type Topic_CRD struct {
	model.CRD `yaml:",inline"`
	Spec      TopicSpec `yaml:"spec"`
}

const (
	apiVer    = "platform.confluent.io/v1beta1"
	kindTopic = "KafkaTopic"
)

func NewKafkaCfkExporter(config config.Config) *KafkaCfkExporter {
	return &KafkaCfkExporter{
		CfkExporter: CfkExporter{
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
			ApiVersion: apiVer,
			Kind:       kindTopic,
			Metadata: model.Metadata{
				Name:      v.Name,
				Namespace: e.Namespace,
			},
		}
		crd := &Topic_CRD{
			CRD: *c,
			Spec: TopicSpec{
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
		err := os.WriteFile(outputPath+"/"+v.Name+".yml", file, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
