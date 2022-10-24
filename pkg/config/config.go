package config

type Config struct {
	Cluster         string `yaml:"cluster" validate:"required"`
	BootstrapServer string `yaml:"bootstrapServer" validate:"required"`
	Environment     string `yaml:"environment" validate:"required"`
	ApiKey          string `yaml:"apiKey" validate:"required"`
	ApiSecret       string `yaml:"apiSecret" validate:"required"`
	CCloudUrl       string `yaml:"ccloudUrl" validate:"required"`
}
