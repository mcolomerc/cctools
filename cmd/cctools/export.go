package cctools

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"
	"mcolomerc/cc-tools/pkg/services"
	"os"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"export-info, cluster-export, confluent-exp, exp"},
	Short:   "Export Cluster Info",
	Long:    ` Command to export cluster information.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Export all the resources: ")
		buildConfig(cmd)
		toolsConfig.Export.Resources = []config.Resource{config.ExportTopics, config.ExportConsumerGroups, config.ExportSchemas}
		runExport(cmd)
	},
}

var topicsCmd = &cobra.Command{
	Use:     "topics",
	Aliases: []string{"topic-info, topic-exp, tpc"},
	Short:   "Export Topics Info",
	Long:    ` Command to export Topics information.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Export Topics information command")
		toolsConfig.Export.Resources = []config.Resource{config.ExportTopics}
		runExport(cmd)
	},
}

var cGroupsCmd = &cobra.Command{
	Use:     "consumer-groups",
	Aliases: []string{"cg, cg-info, cg-exp, cgroup"},
	Short:   "Export Topics Info",
	Long:    ` Command to export Consumer Group information.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Export Consumer Group command")
		toolsConfig.Export.Resources = []config.Resource{config.ExportConsumerGroups}
		runExport(cmd)
	},
}

var schemasCmd = &cobra.Command{
	Use:     "schemas",
	Aliases: []string{"schemas-info, schemas-exp, schema"},
	Short:   "Export Schemas Info",
	Long:    ` Command to export Schemas information.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Export Schemas information command")
		toolsConfig.Export.Resources = []config.Resource{config.ExportSchemas}
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
	builder, err2 := services.NewExportHandler(toolsConfig)
	if err2 != nil {
		log.Error("Error building exporters")
		os.Exit(1)
	}
	builder.BuildExport()
}
