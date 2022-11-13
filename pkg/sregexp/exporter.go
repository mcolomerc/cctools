package sregexp

import (
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
)

type SRegExporter interface {
	export.Exporter
	ExportSubjects(schema []model.SubjectVersion, outputPath string) error
	ExportSchemas(schemas []model.Schema, outputPath string) error
}
