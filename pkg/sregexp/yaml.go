package sregexp

import (
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
)

type SRegYamlExporter struct {
	export.YamlExporter
}

func NewSRegYamlExporter() *SRegYamlExporter {
	return &SRegYamlExporter{
		YamlExporter: export.YamlExporter{},
	}
}

func (e SRegYamlExporter) ExportSubjects(subjects []model.SubjectVersion, outputPath string) error {
	return e.YamlExporter.Export(subjects, outputPath+"_subjects")
}

