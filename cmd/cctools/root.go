package cctools

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/kafka"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var cfgFile string
var toolsConfig config.RuntimeConfig
var kafkaService kafka.KafkaService

var rootCmd = &cobra.Command{
	Use:     "cct",
	Aliases: []string{"cct-info, cct, cct-exp, cctexp"},
	Version: version,
	Short:   "cctools - a simple CLI to manage Clonfluent Cloud",
	Long: `a simple CLI to manage Clonfluent Cloud
    
One can use cctools to ...`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func initConfig() {
	tConfig, err := config.ConfigBuilder{}.Build(cfgFile)
	if err != nil {
		log.Fatalf("Error Reading Config")
	}
	toolsConfig = tConfig
	kafkaService = *kafka.New(toolsConfig)
}

func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.MarkFlagRequired("config")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}

}
