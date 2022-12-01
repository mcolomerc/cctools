package model

import "mcolomerc/cc-tools/pkg/export"

type CRD struct {
	ApiVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   export.Metadata `yaml:"metadata"`
}
