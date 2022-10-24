package cctools

import (
	"fmt"
	"mcolomerc/cc-tools/pkg/ccloud"
	"mcolomerc/cc-tools/pkg/export"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"export-info, cluster-export, confluent-exp, exp"},
	Short:   "Export Confluent Cloud Cluster Info",
	Long:    ` Command to export Confluent Cloud cluster information to a XLSX file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Export Cluster information command \n")
		ccClient, _ := ccloud.New(toolsConfig)
		topics, err := ccClient.GetTopics()
		if err != nil {
			fmt.Printf("client: could not get topics: %s\n", err)
		}
		export.ExportToExcel(topics, toolsConfig)
		/*cgroups, err := ccClient.GetConsumerGroups()
		if err != nil {
			fmt.Printf("client: could not get cgroups: %s\n", err)
		}
		fmt.Print(cgroups) */
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
