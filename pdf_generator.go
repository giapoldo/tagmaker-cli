package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-pdf/fpdf"
)

func (m *model) makeTag(refX float64, tagNumber int, pdf *fpdf.Fpdf) {

	tagHeight, _ := strconv.ParseFloat(m.pVContents.selectedValues["height"], 64)
	tagWidth, _ := strconv.ParseFloat(m.pVContents.selectedValues["width"], 64)

	rowHeight := tagHeight / float64(len(m.tag.tagTable))
	count := 0.0
	nextLine := 0

	for _, row := range m.tag.tagTable {
		pdf.SetX(refX)
		for _, cell := range row {
			count += cell.widthPerUnit

			if count >= 1.0 {
				nextLine = 1
				count = 0
			} else {
				nextLine = 0
			}

			if cell.centered {
				pdf.SetFontStyle(cell.textStyle)
				if cell.isFieldName {
					pdf.CellFormat(tagWidth*cell.widthPerUnit, rowHeight, cell.refHeader, "1", nextLine, "C", false, 0, "")
				} else {
					pdf.CellFormat(tagWidth*cell.widthPerUnit, rowHeight, m.csvData.rows[tagNumber][cell.refHeader], "1", nextLine, "C", false, 0, "")

				}
			} else {
				pdf.SetFontStyle(cell.textStyle)
				if cell.isFieldName {
					pdf.CellFormat(tagWidth*cell.widthPerUnit, rowHeight, cell.refHeader, "1", nextLine, "L", false, 0, "")

				} else {
					pdf.CellFormat(tagWidth*cell.widthPerUnit, rowHeight, m.csvData.rows[tagNumber][cell.refHeader], "1", nextLine, "L", false, 0, "")

				}
			}
		}
		if nextLine == 1 {
			pdf.SetX(refX)
		}
	}
	//padding/margin

	pdf.CellFormat(tagWidth, 2.0, "", "", 1, "C", false, 0, "")
}

func (m *model) pdfGenerator() {

	tagWidth, _ := strconv.ParseFloat(m.pVContents.selectedValues["width"], 64)
	tagHeight, _ := strconv.ParseFloat(m.pVContents.selectedValues["height"], 64)
	fontSize, _ := strconv.ParseFloat(m.pVContents.selectedValues["fontSize"], 64)

	pdf := pdfSetup(m.pVContents.selectedValues["paper"])
	pdf.SetFont("Mechanical", "", fontSize)

	m.currentTag = 0
	columnFillcounter := 0
	w, h := pdf.GetPageSize()
	left, top, right, bottom := pdf.GetMargins()
	x := left
	secondColumnWidth := (w - left - right) - tagWidth - 2

	for range m.csvData.rows {

		if (h-bottom)-pdf.GetY() < tagHeight && columnFillcounter == 0 && secondColumnWidth >= tagWidth {
			x = left + tagWidth + 2
			pdf.SetXY(x, top)
			columnFillcounter = 1
		}
		if (h-bottom)-pdf.GetY() < tagHeight && columnFillcounter == 1 {
			pdf.AddPage()
			x = left
			pdf.SetXY(x, top)
			columnFillcounter = 0
		}

		m.makeTag(x, m.currentTag, pdf)
		m.currentTag++
	}

	err := os.Mkdir("Output", 0777)
	if err != nil {
		log.Printf("Mkdir error: %s", err)
	}

	err = pdf.OutputFileAndClose("Output/Tags.pdf")
	if err != nil {
		log.Printf("PDF output error: %s", err)
	}
}

// func CreateTagTable() {

// }
