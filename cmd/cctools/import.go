package cctools

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Short:   "Import files to a Kafka destination",
	Long:    ` Command to import cluster resources  to another cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Import all the resources: ")
		buildConfig(cmd)
		runCopy(cmd)
	},
}

var importTopicsCmd = &cobra.Command{
	Use:     "topics",
	Aliases: []string{"topic-cp, tpic-cp, tpc"},
	Short:   "Import Topics Info",
	Long:    ` Command to import from source files and create destination Topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Import Topics command")
		toolsConfig.Export.Resources = []config.Resource{config.ExportTopics}
		runCopy(cmd)
	},
}

func init() {
	// Flags
	copyCmd.AddCommand(copyTopicsCmd)
	rootCmd.AddCommand(copyCmd)
}

func runImport(cmd *cobra.Command) {
	log.Info("Importing ...")
}
