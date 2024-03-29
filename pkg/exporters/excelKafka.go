package exporters

import (
	"mcolomerc/cc-tools/pkg/model"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type KafkaExcelExporter struct {
	ParentKafkaExporter
	ExcelExporter
}

func NewKafkaExcelExporter() *KafkaExcelExporter {
	return &KafkaExcelExporter{
		ExcelExporter: ExcelExporter{},
	}
}

func (e KafkaExcelExporter) ExportConsumerGroups(cgroups []model.ConsumerGroup, outputPath string) error {
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("ConsumerGroups")
	f.SetCellValue("ConsumerGroups", "A1", "ConsumerGroupID")
	f.SetCellValue("ConsumerGroups", "B1", "PartitionAssignor")
	f.SetCellValue("ConsumerGroups", "C1", "State")
	f.SetCellValue("ConsumerGroups", "D1", "Consumers")
	f.SetCellValue("ConsumerGroups", "E1", "LagSummary")
	f.SetCellValue("ConsumerGroups", "F1", "Lags")

	for key, value := range cgroups {
		f.SetCellValue("ConsumerGroups", "A"+strconv.Itoa(key+2), value.ConsumerGroupID)
		f.SetCellValue("ConsumerGroups", "B"+strconv.Itoa(key+2), value.PartitionsAssignor)
		f.SetCellValue("ConsumerGroups", "C"+strconv.Itoa(key+2), value.State)
		f.SetCellValue("ConsumerGroups", "D"+strconv.Itoa(key+2), value.Consumers)
		f.SetCellValue("ConsumerGroups", "E"+strconv.Itoa(key+2), value.LagSummary)
		f.SetCellValue("ConsumerGroups", "F"+strconv.Itoa(key+2), value.Lags)
	}

	f.SetActiveSheet(index)
	f.Sheet.Delete("Sheet1")
	// Save spreadsheet by the given path.
	if err := f.SaveAs(outputPath + "_consumer_groups.xlsx"); err != nil {
		return err
	}
	return nil
}

func (e KafkaExcelExporter) ExportTopics(topics []model.Topic, outputPath string) error {
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("Topics")

	f.SetCellValue("Topics", "A1", "Topic")
	f.SetCellValue("Topics", "B1", "Partitions")
	f.SetCellValue("Topics", "C1", "Replication Factor")
	f.SetCellValue("Topics", "D1", "MinISR")
	f.SetCellValue("Topics", "E1", "Retention Time MS")
	f.SetCellValue("Topics", "F1", "Role Bindings")
	f.SetCellValue("Topics", "G1", "Configs")

	for key, value := range topics {
		f.SetCellValue("Topics", "A"+strconv.Itoa(key+2), value.Name)
		f.SetCellValue("Topics", "B"+strconv.Itoa(key+2), value.Partitions)
		f.SetCellValue("Topics", "C"+strconv.Itoa(key+2), value.ReplicationFactor)
		f.SetCellValue("Topics", "D"+strconv.Itoa(key+2), value.MinIsr)
		f.SetCellValue("Topics", "E"+strconv.Itoa(key+2), value.RetentionTime)
		f.SetCellValue("Topics", "F"+strconv.Itoa(key+2), getRoleBindings(value.RoleBindings))
		f.SetCellValue("Topics", "G"+strconv.Itoa(key+2), getConfigs(value.Configs))
	}

	f.SetActiveSheet(index)
	f.Sheet.Delete("Sheet1")
	// Save spreadsheet by the given path.
	if err := f.SaveAs(outputPath + "_topics.xlsx"); err != nil {
		return err
	}
	return nil
}

func getRoleBindings(roleBindings []model.RoleBinding) string {
	var roleBindingsCell string

	for _, value := range roleBindings {
		roleBindingsCell = roleBindingsCell + value.RoleName + "=" + value.Users.(string) + "\n"
	}

	return roleBindingsCell
}

func getConfigs(configs []model.TopicConfig) string {
	var configCell string

	for _, value := range configs {
		configCell = configCell + value.Name + "=" + value.Value.(string) + "\n"
	}

	return configCell
}
