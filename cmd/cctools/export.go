package cctools

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/log"
	"os"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export Cluster metadata",
	Long:  ` Command to export cluster metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Export all the resources: ")
		//validate args
		if len(args) > 0 {
			if args[0] != config.Topic.String() && args[0] != config.ConsumerGroup.String() && args[0] != config.Schema.String() {
				log.Error("Invalid resource to export: " + args[0])
				log.Error("Valid resources: " + config.Topic.String() + ", " + config.ConsumerGroup.String() + ", " + config.Schema.String())
				cmd.Help()
				os.Exit(1)
			}
		}
		buildConfig(cmd)
		toolsConfig.Export.Resources = []config.Resource{config.Topic, config.ConsumerGroup, config.Schema}
		runExport(cmd)
	},
}

var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "Export Topics",
	Long:  ` Command to export Topics metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Export Topics information command")
		toolsConfig.Export.Resources = []config.Resource{config.Topic}
		runExport(cmd)
	},
}

var cGroupsCmd = &cobra.Command{
	Use:   "consumer-groups",
	Short: "Export Consumer Groups",
	Long:  ` Command to export Consumer Group information.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Export Consumer Group command")
		toolsConfig.Export.Resources = []config.Resource{config.ConsumerGroup}
		runExport(cmd)
	},
}

var schemasCmd = &cobra.Command{
	Use:   "schemas",
	Short: "Export Schemas Info",
	Long:  ` Command to export Schemas information.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Export Schemas information command")
		toolsConfig.Export.Resources = []config.Resource{config.Schema}
		runExport(cmd)
	},
}

func init() {
	// Flags
	exportCmd.PersistentFlags().StringP("output", "o", "", "Output format. Possible values: json, yaml, hcl, cfk, clink")
	exportCmd.MarkPersistentFlagRequired("output")
	exportCmd.AddCommand(topicsCmd)
	exportCmd.AddCommand(cGroupsCmd)
	exportCmd.AddCommand(schemasCmd)
	rootCmd.AddCommand(exportCmd)
}

func buildConfig(cmd *cobra.Command) {
	cfgFile, err := cmd.Flags().GetString("config")
	if err != nil {
		log.Info("Error getting config flag")
		os.Exit(1)
	}
	tConfig, err := config.ConfigBuilder{}.Build(cfgFile)
	if err != nil {
		log.Error("Error Reading Config")
		os.Exit(1)
	}
	toolsConfig = tConfig
}
func runExport(cmd *cobra.Command) {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Error("Error getting output flag")
		log.Error(err)
		os.Exit(1)
	}
	log.Info("Format: " + output) //TODO: Add multiple exporters to the array.
	toolsConfig.Export.Exporters = []config.Exporter{config.Exporter(output)}
	builder, err2 := export.NewExportHandler(toolsConfig)
	if err2 != nil {
		log.Error("Error building exporters")
		os.Exit(1)
	}
	builder.BuildExport()
}
