package cctools

import (
	"fmt"
	"mcolomerc/cc-tools/pkg/export"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"export-info, cluster-export, confluent-exp, exp"},
	Short:   "Export Confluent Cloud Cluster Info",
	Long:    ` Command to export Confluent Cloud cluster information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Export Cluster information command \n")
		exportResult, err := kafkaService.Export()
		if err != nil {
			fmt.Printf("client: could not get topics: %s\n", err)
		}

		exportExecutors := toolsConfig.Exporters
		outputPath := toolsConfig.UserConfig.Export.Output + "/" + toolsConfig.UserConfig.Cluster
		done := make(chan bool, len(exportExecutors))
		for _, v := range exportExecutors {
			go func(v export.Exporter) {
				err := v.ExportTopics(exportResult.Topics, outputPath)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				done <- true
			}(v)
		}
		for i := 0; i < len(exportExecutors); i++ {
			<-done
		}
		close(done)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
