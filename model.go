package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput           textinput.Model
	textInputVisibility bool
	textValue           string
	inputCaller         string
	flexBox             *flexbox.FlexBox
	tag                 Tag
	currCursorRow       int
	currCursorCell      int
}

func (m *model) createRows(text string) {

	rows := []*flexbox.Row{}

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	rows = append(rows, firstRow)

	// Add tag rows
	for _, row := range m.tag.table {
		_fbRow := m.flexBox.NewRow()

		if _fbRow == nil {
			panic("could not find the table row")
		}
		// Add first padding cell before adding content cells
		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		if len(row) == 1 {
			row[0].widthPerUnit = 1.0
		}

		// Add content cells
		for j, cell := range row {
			style := m.cellStyleSelector(cell, styleNormal)

			_fbRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style))
			_cell := _fbRow.GetCell(j + 1).SetContent(cell.text) // +1 because of cell padding
			if _cell == nil {
				panic("could not find the table cell")
			}
		}
		// Add closing padding cell
		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, _fbRow)
	}
	// Add closing padding row
	lastRow := m.flexBox.NewRow()
	lastRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG).SetContent(text)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	if m.textInputVisibility {
		lastRow.GetCell(1).SetStyle(styleTextInput)
	}
	rows = append(rows, lastRow)

	m.flexBox.SetRows(rows)

	// Highlight the current content row and cell as selected
	if (m.currCursorRow >= 0 && m.currCursorRow < len(m.tag.table)) &&
		(m.currCursorCell >= 0 && m.currCursorCell < len(m.tag.table[m.currCursorRow])) {

		cell := m.tag.table[m.currCursorRow][m.currCursorCell]
		style := m.cellStyleSelector(cell, styleSelected)
		rows[m.FBCursorRow()].GetCell(m.FBCursorCell()).SetStyle(style)
	}

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling CreateRows, we always want to overwrite to refresh the view.
}

func InitialModel() *model {

	dm := model{}
	dm.flexBox = flexbox.New(0, 0)
	// dm.flexBox.LockRowHeight(3)
	dm.tag = Tag{
		// width:  80.0,
		// height: 40.0,
		table: TagTable{
			{
				{widthPerUnit: 1.0,
					text:      "Collection",
					centered:  true,
					textStyle: "B"},
			},
			{
				{widthPerUnit: 1.0,
					text:      "Milestone",
					centered:  true,
					textStyle: ""},
			}, {
				{widthPerUnit: 1.0,
					text:      "",
					centered:  false,
					textStyle: ""},
			},
			{
				{widthPerUnit: 0.5,
					text:      "Field 1",
					centered:  false,
					textStyle: "B"},
				{widthPerUnit: 0.5,
					text:      "UTF8 1",
					centered:  false,
					textStyle: ""},
			},
			{
				{widthPerUnit: 0.3,
					text:      "Nombre",
					centered:  false,
					textStyle: "B"},
				{widthPerUnit: 0.2,
					text:      "GMG",
					centered:  false,
					textStyle: ""},
				{widthPerUnit: 0.3,
					text:      "Fecha",
					centered:  false,
					textStyle: "BI"},
				{widthPerUnit: 0.2,
					text:      "2024",
					centered:  false,
					textStyle: "I"},
			},
		},
	}

	// Cursors start at -1 to avoid starting with a cell selected
	dm.currCursorRow = 0
	dm.currCursorCell = 0

	dm.createRows("")

	return &dm
}

func (m *model) cellStyleSelector(cell Cell, baseStyle lipgloss.Style) (style lipgloss.Style) {

	style = baseStyle

	if !cell.centered && cell.textStyle == "B" {
		style = baseStyle.Bold(true)

	} else if !cell.centered && cell.textStyle == "I" {
		style = baseStyle.Italic(true)

	} else if !cell.centered && cell.textStyle == "BI" {
		style = baseStyle.Bold(true).Italic(true)

	} else if cell.centered && cell.textStyle == "" {
		style = baseStyle.AlignHorizontal(lipgloss.Center)

	} else if cell.centered && cell.textStyle == "B" {
		style = baseStyle.Bold(true).AlignHorizontal(lipgloss.Center)

	} else if cell.centered && cell.textStyle == "I" {
		style = baseStyle.Italic(true).AlignHorizontal(lipgloss.Center)

	} else if cell.centered && cell.textStyle == "BI" {
		style = baseStyle.Bold(true).Italic(true).AlignHorizontal(lipgloss.Center)

	}

	return // default baseStyle is always regular font, left aligned and vertical centered
}
