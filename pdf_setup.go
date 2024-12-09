package main

import (
	"path/filepath"

	"github.com/go-pdf/fpdf"
)

type pageSize struct {
	width  int
	height int
}

var pageSizes map[string]pageSize

func pdfSetup(pageSize string) *fpdf.Fpdf {
	project_dir := "."
	fontPath := filepath.Join(project_dir, "font")
	pdf := fpdf.New("P", "mm", pageSize, fontPath) // 1cm margins by default. TODO args need to connect to GUI/TUI
	pdf.AddPage()
	pdf.AddUTF8Font("Trueno", "", "Trueno.ttf")
	pdf.AddUTF8Font("Trueno", "I", "Trueno-I.ttf")
	pdf.AddUTF8Font("Trueno", "B", "Trueno-B.ttf")
	pdf.AddUTF8Font("Trueno", "BI", "Trueno-BI.ttf")
	pdf.AddUTF8Font("Mechanical", "", "MechanicalCondensed.ttf")
	pdf.AddUTF8Font("Mechanical", "I", "MechanicalCondensed-I.ttf")
	pdf.AddUTF8Font("Mechanical", "B", "MechanicalCondensed-B.ttf")
	pdf.AddUTF8Font("Mechanical", "BI", "MechanicalCondensed-BI.ttf")
	pdf.AddUTF8Font("SourceSerif4", "", "SourceSerif4.ttf")
	pdf.AddUTF8Font("SourceSerif4", "I", "SourceSerif4-I.ttf")
	pdf.AddUTF8Font("SourceSerif4", "B", "SourceSerif4-B.ttf")
	pdf.AddUTF8Font("SourceSerif4", "BI", "SourceSerif4-BI.ttf")

	return pdf

}
