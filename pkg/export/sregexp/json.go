package sregexp

import (
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
)

type SRegJsonExporter struct {
	export.JsonExporter
}

func NewSRegJsonExporter() *SRegJsonExporter {
	return &SRegJsonExporter{
		JsonExporter: export.JsonExporter{},
	}
}

func (e SRegJsonExporter) ExportSubjects(subjects []model.SubjectVersion, outputPath string) error {
	return e.JsonExporter.Export(subjects, outputPath+"_subjects")
}
