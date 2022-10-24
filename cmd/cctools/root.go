package cctools

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/config"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "0.0.1"
var cfgFile string
var toolsConfig config.Config

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
