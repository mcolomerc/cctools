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
	Short:   "cctools - a simple CLI to manage migrations",
	Long: `a simple CLI to manage migrations,
    
One can use cctools to ...
 - Export Topics and ACls from a Kafka Source cluster to different formats (JSON,YML,CFK,HCL, ...)
 - Export Schemas and Subjects from Schema Registry to different formats (JSON,YML,CFK,HCL, ...)`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.config.yaml)")
	rootCmd.MarkPersistentFlagRequired("config")
	//cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
