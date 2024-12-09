package main

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

func (m *model) makeTag(refX float64, tagNumber int, pdf *fpdf.Fpdf) {

	rowHeight := m.tag.height / float64(len(m.tag.tagTable))
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
					pdf.CellFormat(m.tag.width*cell.widthPerUnit, rowHeight, cell.refHeader, "1", nextLine, "C", false, 0, "")
				} else {
					pdf.CellFormat(m.tag.width*cell.widthPerUnit, rowHeight, m.csvData.rows[tagNumber][cell.refHeader], "1", nextLine, "C", false, 0, "")

				}
			} else {
				pdf.SetFontStyle(cell.textStyle)
				if cell.isFieldName {
					pdf.CellFormat(m.tag.width*cell.widthPerUnit, rowHeight, cell.refHeader, "1", nextLine, "L", false, 0, "")

				} else {
					pdf.CellFormat(m.tag.width*cell.widthPerUnit, rowHeight, m.csvData.rows[tagNumber][cell.refHeader], "1", nextLine, "L", false, 0, "")

				}
			}
		}
		if nextLine == 1 {
			pdf.SetX(refX)
		}
	}
	//padding/margin

	pdf.CellFormat(m.tag.width, 2.0, "", "", 1, "C", false, 0, "")
}

func (m *model) pdfGenerator() {

	pdf := pdfSetup(m.paperSize)
	pdf.SetFont("Mechanical", "", m.tag.fontSize)

	m.currentTag = 0
	columnFillcounter := 0
	w, h := pdf.GetPageSize()
	left, top, right, bottom := pdf.GetMargins()
	x := left
	secondColumnWidth := (w - left - right) - m.tag.width - 2

	for range m.csvData.rows {

		if (h-bottom)-pdf.GetY() < m.tag.height && columnFillcounter == 0 && secondColumnWidth >= m.tag.width {
			x = left + m.tag.width + 2
			pdf.SetXY(x, top)
			columnFillcounter = 1
		}
		if (h-bottom)-pdf.GetY() < m.tag.height && columnFillcounter == 1 {
			pdf.AddPage()
			x = left
			pdf.SetXY(x, top)
			columnFillcounter = 0
		}

		m.makeTag(x, m.currentTag, pdf)
		m.currentTag++
	}

	err := pdf.OutputFileAndClose("Output/Tags.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

// func CreateTagTable() {

// }
