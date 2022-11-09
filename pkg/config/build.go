package config

import (
	"fmt"
	"log"

	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigBuilder struct {
	Config Config
}

func (c ConfigBuilder) Build(cfgFile string) (Config, error) {

	// Read config file with Viper
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
	// Validate Config
	if err := viper.Unmarshal(&c.Config); err != nil {
		log.Printf("unable to unmarshall the config %v", err)
		return c.Config, err
	}
	validate := validator.New()
	if err := validate.Struct(&c.Config); err != nil {
		log.Printf("Missing required attributes %v\n", err)
		return c.Config, err
	}
	//Build Output
	if err := c.buildOutput(); err != nil {
		log.Printf("Can't mount Output folder %v\n", err)
		return c.Config, err
	}

	return c.Config, nil
}

// Build output folder
func (c ConfigBuilder) buildOutput() error {
	if _, err := os.Stat(c.Config.Export.Output); os.IsNotExist(err) {
		log.Printf("Export output directory: %s not found. Creating...", c.Config.Export.Output)
		err := os.Mkdir(c.Config.Export.Output, os.ModePerm)
		if err != nil {
			log.Fatalf("Export output directory: %s - %v", c.Config.Export.Output, err)
			return err
		}
	}
	return nil
}
