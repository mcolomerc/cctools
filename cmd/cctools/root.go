package cctools

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/export"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "0.0.1"
var cfgFile string
var toolsConfig config.Config

var exportExecutors []export.Export

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
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&toolsConfig); err != nil {
		log.Fatalf("unable to unmarshall the config %v", err)
	}
	validate := validator.New()
	if err := validate.Struct(&toolsConfig); err != nil {
		log.Fatalf("Missing required attributes %v\n", err)
	}

	for _, v := range toolsConfig.Export.Exporters {
		if v == config.Excel {
			exportExecutors = append(exportExecutors, &export.ExcelExporter{})
		}
		if v == config.Json {
			exportExecutors = append(exportExecutors, &export.JsonExporter{})
		}
		if v == config.Yaml {
			exportExecutors = append(exportExecutors, &export.YamlExporter{})
		}
	}

	if _, err := os.Stat(toolsConfig.Export.Output); os.IsNotExist(err) {
		log.Printf("Export output directory: %s not found. Creating...", toolsConfig.Export.Output)
		err := os.Mkdir(toolsConfig.Export.Output, os.ModePerm)
		if err != nil {
			log.Fatalf("Export output directory: %s - %v", toolsConfig.Export.Output, err)
			os.Exit(1)
		}
	}

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
