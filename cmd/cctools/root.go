package cctools

import (
	"mcolomerc/cc-tools/pkg/config"
	log "mcolomerc/cc-tools/pkg/log"

	"mcolomerc/cc-tools/pkg/services"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var cfgFile string
var toolsConfig config.Config
var exportHandler services.ExportHandler

var rootCmd = &cobra.Command{
	Use:     "cctools",
	Aliases: []string{"cct-info, cct, cct-exp, cctexp"},
	Version: version,
	Short:   "cctools - a simple CLI to manage Apache Kafka migrations",
	Long:    `a simple CLI to manage Apache Kafka migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Run cctools without command.")
		cmd.Help()
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file")
	rootCmd.MarkPersistentFlagRequired("config")
	//cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
