package cctools

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/imprt"
	"mcolomerc/cc-tools/pkg/log"
	"os"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import files to a Kafka destination",
	Long:  ` Command to import cluster resources  to another cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Import all the resources: ")
		if len(args) > 0 {
			if args[0] != config.Topic.String() && args[0] != config.Schema.String() {
				log.Error("Invalid resource to export: " + args[0])
				log.Error("Valid resources: " + config.Topic.String() + " ," + config.Schema.String())
				cmd.Help()
				os.Exit(1)
			}
		}
		buildConfig(cmd)
		toolsConfig.Import.Resources = []config.Resource{config.Topic, config.Schema}
		runCopy(cmd)
	},
}

var importTopicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "Import Topics Info",
	Long:  ` Command to import from source files and create destination Topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Import Topics command")
		toolsConfig.Import.Resources = []config.Resource{config.Topic}
		runImport(cmd)
	},
}

var importSchemasCmd = &cobra.Command{
	Use:   "schemas",
	Short: "Import Schemas ",
	Long:  ` Command to import from source files and create destination Schemas.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Import Schemas command")
		toolsConfig.Import.Resources = []config.Resource{config.Schema}
		runImport(cmd)
	},
}

func init() {
	// Flags
	importCmd.AddCommand(importTopicsCmd)
	importCmd.AddCommand(importSchemasCmd)
	rootCmd.AddCommand(importCmd)
}

func runImport(cmd *cobra.Command) {
	log.Info("Importing ...")
	importService, err := imprt.NewImportHandler(toolsConfig)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	importService.Import()
}
