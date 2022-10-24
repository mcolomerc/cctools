package export

import (
	"fmt"
	"mcolomerc/cc-tools/pkg/ccloud"
	"mcolomerc/cc-tools/pkg/config"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ExportToExcel(topics []ccloud.Topic, config config.Config) {
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("Topics")

	f.SetCellValue("Topics", "A1", "Topic")
	f.SetCellValue("Topics", "B1", "Partitions")
	f.SetCellValue("Topics", "C1", "Replication Factor")
	f.SetCellValue("Topics", "D1", "Configs")

	for key, value := range topics {
		f.SetCellValue("Topics", "A"+strconv.Itoa(key+2), value.Name)
		f.SetCellValue("Topics", "B"+strconv.Itoa(key+2), value.Partitions)
		f.SetCellValue("Topics", "C"+strconv.Itoa(key+2), value.ReplicationFactor)
		f.SetCellValue("Topics", "D"+strconv.Itoa(key+2), getConfigs(value.Configs))
	}

	f.SetActiveSheet(index)
	f.Sheet.Delete("Sheet1")
	// Save spreadsheet by the given path.
	if err := f.SaveAs(config.Cluster + "_Topics.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func getConfigs(configs []ccloud.TopicConfig) string {
	var configCell string

	for _, value := range configs {
		configCell = configCell + value.Name + "=" + value.Value.(string) + "\n"
	}

	return configCell
}
