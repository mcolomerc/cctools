package sregexp

import (
	"github.com/xuri/excelize/v2"
	"mcolomerc/cc-tools/pkg/export"
	"mcolomerc/cc-tools/pkg/model"
	"strconv"
)

type SRegExcelExporter struct {
	*export.ExcelExporter
}

func NewSRegExcelExporter() *SRegExcelExporter {
	return &SRegExcelExporter{
		ExcelExporter: &export.ExcelExporter{},
	}
}

func (e SRegExcelExporter) ExportSubjects(subjects []model.SubjectVersion, outputPath string) error {
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("Schemas")
	f.SetCellValue("Schemas", "A1", "SchemaId")
	f.SetCellValue("Schemas", "B1", "Subject Name")
	f.SetCellValue("Schemas", "C1", "Version")
	f.SetCellValue("Schemas", "D1", "Schema")

	for key, value := range subjects {
		f.SetCellValue("Schemas", "A"+strconv.Itoa(key+2), value.Id)
		f.SetCellValue("Schemas", "B"+strconv.Itoa(key+2), value.Subject)
		f.SetCellValue("Schemas", "C"+strconv.Itoa(key+2), value.Version)
		f.SetCellValue("Schemas", "D"+strconv.Itoa(key+2), value.Schema)
	}

	f.SetActiveSheet(index)
	f.Sheet.Delete("Sheet1")
	// Save spreadsheet by the given path.
	if err := f.SaveAs(outputPath + "_subjects.xlsx"); err != nil {
		return err
	}
	return nil
}
