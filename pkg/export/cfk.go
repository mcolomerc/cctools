package export

import (
	"io/ioutil"
	"mcolomerc/cc-tools/pkg/model"

	"gopkg.in/yaml.v2"
)

type CfkExporter struct {
	Namespace      string
	KafkaRestClass string
}

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
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
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

const (
	apiVersion = "platform.confluent.io/v1beta1"
	kind       = "KafkaTopic"
)

func (e CfkExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	for _, v := range topics {
		configs := make(map[string]interface{})
		for _, vc := range v.Configs {
			configs[vc.Name] = vc.Value
		}
		crd := &Topic_CRD{
			ApiVersion: apiVersion,
			Kind:       kind,
			Metadata: Metadata{
				Name:      v.Name,
				Namespace: e.Namespace,
			},
			Spec: Spec{
				Replicas:   v.ReplicationFactor,
				Partitions: v.Partitions,
				Configs:    configs,
				KafkaRestClass: KafkaRestClass{
					Name: e.KafkaRestClass,
				},
			},
		}
		file, errJson := yaml.Marshal(crd)
		if errJson != nil {
			return errJson
		}
		err := ioutil.WriteFile(outputPath+"_topic_"+v.Name+".yml", file, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e CfkExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	return nil
}
