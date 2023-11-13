package exporters

import (
	"mcolomerc/cc-tools/pkg/model"
)

type SRegExporter interface {
	Exporter
	ExportSubjects(schema []model.SubjectVersion, outputPath string) error
	ExportSchemas(schemas []model.Schema, outputPath string) error
}
