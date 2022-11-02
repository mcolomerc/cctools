package config

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestConfig(t *testing.T) {

	configFile := "./testdata/config.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err != nil {
		t.Fatalf(`&ConfigBuilder{}.Build`)
	}
	want := "superUser"

	if config.UserConfig.Credentials.Key != want || err != nil {
		t.Fatalf(`TestConfig config.UserConfig.Credentials.Key %s = %s, nil`, config.UserConfig.Credentials.Key, want)
	}
}

func TestConfigCloud(t *testing.T) {

	configFile := "./testdata/config_cc.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err != nil {
		t.Fatalf(`&ConfigBuilder{}.Build`)
	}
	empty := config.UserConfig.Credentials.Certificates != Certificates{}

	if empty {
		t.Fatalf(`Tconfig.UserConfig.Credentials.Certificates %v, %v`, config.UserConfig.Credentials.Certificates, err)
	}
}

func TestConfigYaml(t *testing.T) {

	configFile := "./testdata/config_exp_yaml.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err != nil {
		t.Fatalf(`&ConfigBuilder{}.Build`)
	}
	for _, exp := range config.UserConfig.Export.Exporters {
		if exp != Yaml {
			t.Fatalf(`Got %v, Expected %v`, exp, Yaml)
		}
	}
}

func TestConfigJson(t *testing.T) {

	configFile := "./testdata/config_exp_json.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err != nil {
		t.Fatalf(`&ConfigBuilder{}.Build`)
	}

	for _, exp := range config.UserConfig.Export.Exporters {
		if exp != Json {
			t.Fatalf(`Got %v, Expected %v`, exp, Yaml)
		}
	}
}
func TestConfigMissingExpOutPut(t *testing.T) {

	configFile := "./testdata/config_exp_missing_output.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err == nil {
		t.Fatalf(`Output should be required %v`, config.UserConfig.Export.Output)
	}

}
