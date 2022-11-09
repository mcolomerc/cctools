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

	if config.Credentials.Key != want || err != nil {
		t.Fatalf(`TestConfig config.Credentials.Key %s = %s, nil`, config.Credentials.Key, want)
	}
}

func TestConfigCloud(t *testing.T) {

	configFile := "./testdata/config_cc.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err != nil {
		t.Fatalf(`&ConfigBuilder{}.Build`)
	}
	empty := config.Credentials.Certificates != Certificates{}

	if empty {
		t.Fatalf(`Tconfig.Credentials.Certificates %v, %v`, config.Credentials.Certificates, err)
	}
}

func TestConfigYaml(t *testing.T) {

	configFile := "./testdata/config_exp_yaml.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err != nil {
		t.Fatalf(`&ConfigBuilder{}.Build`)
	}
	for _, exp := range config.Export.Exporters {
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

	for _, exp := range config.Export.Exporters {
		if exp != Json {
			t.Fatalf(`Got %v, Expected %v`, exp, Yaml)
		}
	}
}
func TestConfigMissingExpOutPut(t *testing.T) {

	configFile := "./testdata/config_exp_missing_output.yml"

	config, err := ConfigBuilder{}.Build(configFile)
	if err == nil {
		t.Fatalf(`Output should be required %v`, config.Export.Output)
	}

}
