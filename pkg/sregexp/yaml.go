package sregexp

import (
	"fmt"
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
	done := make(chan bool, len(subjects))
	for _, s := range subjects {
		go func(subject model.SubjectVersion, outputPath string) {
			out := fmt.Sprintf("%s%s_%d", outputPath, subject.Subject, subject.Version)
			e.YamlExporter.Export(subject, out)
			done <- true
		}(s, outputPath)
	}
	for i := 0; i < len(subjects); i++ {
		<-done
	}
	close(done)
	return nil
}

func (e SRegYamlExporter) ExportSchemas(schemas []model.Schema, outputPath string) error {
	done := make(chan bool, len(schemas))
	for _, s := range schemas {
		go func(schema model.Schema, outputPath string) {
			out := fmt.Sprintf("%s%s_%d", outputPath, schema.Subject, schema.Version)
			e.YamlExporter.Export(schema, out)
			done <- true
		}(s, outputPath)
	}
	for i := 0; i < len(schemas); i++ {
		<-done
	}
	close(done)
	return nil
}
