package main

import (
	"github.com/76creates/stickers/flexbox"
)

func (m *model) TagView(text string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(5)

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

	if m.withTextInput {
		lastRow.GetCell(1).SetStyle(styleTextInput)
	}
	rows = append(rows, lastRow)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).
		SetContent("\nArrows to move\nA: Insert row below\t\tZ: Delete current row\t\tX: Delete current cell\nS: Insert cell to the left\tD: Insert cell to the right\tB: Bind CSV data to Tag fields"))

	rows = append(rows, helpRow)

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling TagView, we always want to overwrite to refresh the view.
	m.flexBox.SetRows(rows)

	// Highlight the current content row and cell as selected
	if (m.tagCursorRow >= 0 && m.tagCursorRow < len(m.tag.table)) &&
		(m.tagCursorCell >= 0 && m.tagCursorCell < len(m.tag.table[m.tagCursorRow])) {

		cell := m.tag.table[m.tagCursorRow][m.tagCursorCell]
		style := m.cellStyleSelector(cell, styleSelected)
		rows[m.FBCursorRow()].GetCell(m.FBCursorCell()).SetStyle(style)
	}
}

func (m *model) WelcomeView() {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(12)

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	rows = append(rows, firstRow)

	textRow := m.flexBox.NewRow()
	textRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	welcomeText := "Welcome to TagBuilder.\n\nThis tool helps you build a tag format for identification of elements in collections, samples or fieldwork findings around your tabulated data.\n\nInitially for museological work, but use it for whatever yo want.\n\nPlease see the provided CSV table for the format.\n\nPress Ctrl+C or Q to quit or any other key to continue"

	textRow.AddCells(flexbox.NewCell(100, 1).SetStyle(styleWelcome).SetContent(welcomeText))
	textRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	rows = append(rows, textRow)

	// Add closing padding row
	lastRow := m.flexBox.NewRow()
	lastRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	rows = append(rows, lastRow)

	m.flexBox.SetRows(rows)

}
