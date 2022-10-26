package export

import (
	"mcolomerc/cc-tools/pkg/ccloud"
	"mcolomerc/cc-tools/pkg/config"
)

type Export interface {
	ExportTopics(topics []ccloud.Topic, config config.Config) error
}
