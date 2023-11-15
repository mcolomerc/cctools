package cctools

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/copy"
	"mcolomerc/cc-tools/pkg/log"
	"os"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy",
	Aliases: []string{"copy-info, cluster-copy, confluent-copy, cp"},
	Short:   "Copy source Kafka cluster resource to a Kafka destination",
	Long:    ` Command to copy cluster resources  to another cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Copy all the resources: ")
		//validate args
		if len(args) > 0 {
			if args[0] != config.Topic.String() {
				log.Error("Invalid resource to export: " + args[0])
				log.Error("Valid resources: " + config.Topic.String())
				cmd.Help()
				os.Exit(1)
			}
		}
		buildConfig(cmd)
		runCopy(cmd)
	},
}

var copyTopicsCmd = &cobra.Command{
	Use:     "topics",
	Aliases: []string{"topic-cp, tpic-cp, tpc"},
	Short:   "Copy Topics Info",
	Long:    ` Command to copy from source Kafka and create destination Topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildConfig(cmd)
		log.Info("Copy Topics command")
		toolsConfig.Export.Resources = []config.Resource{config.Topic}
		runCopy(cmd)
	},
}

func init() {
	// Flags
	copyCmd.AddCommand(copyTopicsCmd)
	rootCmd.AddCommand(copyCmd)
}

func runCopy(cmd *cobra.Command) {
	kService, err := copy.NewCopyHandler(toolsConfig)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	kService.Copy()
}
