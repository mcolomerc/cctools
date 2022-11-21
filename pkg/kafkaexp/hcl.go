package kafkaexp

import (
	"fmt"
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/model"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type HclExporter struct {
	ParentKafkaExporter
	export.HclExporter
}

type CtyTopics struct {
	Environment    string `cty:"environment"`
	Cluster        string `cty:"cluster"`
	RBAC           string `cty:"rbac_enabled"`
	ServiceAccount struct {
		Name string `cty:"name"`
		Role string `cty:"role"`
	} `cty:"serv_account"`
	CtyTopics []CtyTopic `cty:"topics"`
}
type CtyTopic struct {
	Name       string            `cty:"name"`
	Partitions int               `cty:"partitions"`
	Config     map[string]string `cty:"config"`
}

func NewHclExporter() *HclExporter {
	return &HclExporter{
		HclExporter: export.HclExporter{},
	}
}

func (e HclExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	log.Debug("Exporting topics to HCL")

	f := hclwrite.NewEmptyFile()

	// create new file on system
	tfFile, err := os.Create(outputPath + "topics.tf")
	if err != nil {
		fmt.Println(err)
	}

	rootBody := f.Body()

	rootBody.SetAttributeValue("environment", cty.StringVal("<ENV_ID>"))
	rootBody.AppendNewline()
	rootBody.SetAttributeValue("cluster", cty.StringVal("<CLUSTER_ID>"))
	rootBody.AppendNewline()
	rootBody.SetAttributeValue("rbac_enabled", cty.False)
	rootBody.AppendNewline()
	rootBody.SetAttributeValue("serv_account", cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal("<SERVICE_ACCOUNT>"),
		"role": cty.StringVal("CloudClusterAdmin"),
	}))
	rootBody.AppendNewline()
	var vals []cty.Value
	for _, topic := range topics {
		parts := topic.Partitions.(float64)
		objTopic := cty.ObjectVal(map[string]cty.Value{
			"name":       cty.StringVal(topic.Name),
			"partitions": cty.NumberIntVal(int64(parts)),
			"config":     cty.MapVal(buildConfigs(topic)),
		})
		vals = append(vals, objTopic)
	}
	rootBody.SetAttributeValue("topics", cty.ListVal(vals))
	tfFile.Write(f.Bytes())

	return nil
}

func buildConfigs(topic model.Topic) map[string]cty.Value {
	res := make(map[string]cty.Value)
	for _, config := range topic.Configs {
		res[config.Name] = cty.StringVal(config.Value.(string))
	}
	return res
}
