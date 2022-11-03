package cctools

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"export-info, cluster-export, confluent-exp, exp"},
	Short:   "Export Confluent Cloud Cluster Info",
	Long:    ` Command to export Confluent Cloud cluster information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Export Cluster information command \n")
		exportHandler.BuildExport()
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
