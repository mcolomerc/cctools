package config

import (
	"fmt"
	"log"
	"mcolomerc/cc-tools/pkg/export"
	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RuntimeConfig struct {
	UserConfig Config
	Exporters  []export.Exporter
}
type ConfigBuilder struct {
	RuntimeConfig RuntimeConfig
}

func (c ConfigBuilder) Build(cfgFile string) (RuntimeConfig, error) {

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
	// Validate UserConfig
	if err := viper.Unmarshal(&c.RuntimeConfig.UserConfig); err != nil {
		log.Printf("unable to unmarshall the config %v", err)
		return c.RuntimeConfig, err
	}
	validate := validator.New()
	if err := validate.Struct(&c.RuntimeConfig.UserConfig); err != nil {
		log.Printf("Missing required attributes %v\n", err)
		return c.RuntimeConfig, err
	}
	//Build Output
	if err := c.buildOutput(); err != nil {
		log.Printf("Can't mount Output folder %v\n", err)
		return c.RuntimeConfig, err
	}
	//Build Exporters
	c.RuntimeConfig.Exporters = c.buildExporters()
	return c.RuntimeConfig, nil
}

// Build output folder
func (c ConfigBuilder) buildOutput() error {
	if _, err := os.Stat(c.RuntimeConfig.UserConfig.Export.Output); os.IsNotExist(err) {
		log.Printf("Export output directory: %s not found. Creating...", c.RuntimeConfig.UserConfig.Export.Output)
		err := os.Mkdir(c.RuntimeConfig.UserConfig.Export.Output, os.ModePerm)
		if err != nil {
			log.Fatalf("Export output directory: %s - %v", c.RuntimeConfig.UserConfig.Export.Output, err)
			return err
		}
	}
	return nil
}

// Build exporters
func (c ConfigBuilder) buildExporters() []export.Exporter {
	var exporters []export.Exporter
	for _, v := range c.RuntimeConfig.UserConfig.Export.Exporters {
		if v == Excel {
			exporters = append(exporters, &export.ExcelExporter{})
		} else if v == Json {
			exporters = append(exporters, &export.JsonExporter{})
		} else if v == Yaml {
			exporters = append(exporters, &export.YamlExporter{})
		} else {
			fmt.Printf("Unrecognized exporter: %v", v)
		}
	}
	return exporters
}
