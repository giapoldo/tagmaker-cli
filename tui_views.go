package main

import (
	"github.com/76creates/stickers/flexbox"
)

func (m *model) welcome1View() {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(12)

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	rows = append(rows, firstRow)

	textRow := m.flexBox.NewRow()
	textRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	welcomeText := "Welcome to TagBuilder.\n\nThis tool helps you build a tag format for classification of elements in collections, samples or fieldwork findings around your tabulated data.\n\nInitially for museological work, but use it for whatever yo want.\n\nCheck the provided CSV table in the \"Input\" folder for the format.\n\nPress Ctrl+C or Q to quit or any other key to continue"

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

func (m *model) welcome2View() {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(12)

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	rows = append(rows, firstRow)

	textRow := m.flexBox.NewRow()
	textRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	welcomeText := "Press A to add your first row.\nthen use the help below to construct the rest of the tag."

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

func (m *model) tagBuilderView(footerText string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBindList).SetContent("Tag Builder")).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	rows = append(rows, firstRow)

	/* ---------------------------------------------------------------------------*/
	/* ---------------------------------------------------------------------------*/

	// Add tag rows
	for _, row := range m.tag.tagTable {

		dataRow := m.flexBox.NewRow()

		if dataRow == nil {
			panic("could not find the table row")
		}
		// Add first padding tagCellbefore adding content cells
		dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

		// Ensures addition or deletion truncation from float to int ends up at length 1.0
		if len(row) == 1 {
			row[0].widthPerUnit = 1.0
		}

		// Add content cells
		for _, cell := range row {

			style := m.getCellTextStyle(cell, styleNormal)
			if len(cell.refHeader) > 0 {
				if cell.isFieldName {
					dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style).
						SetContent(cell.refHeader))
				} else {
					dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style).
						SetContent(m.csvData.rows[m.currentTag][cell.refHeader]))
				}

			} else {
				dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style))
			}
		}

		// Add closing padding cell
		dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, dataRow)
	}

	// // Add tag rows
	// for _, row := range m.tag.tagTable {

	// 	dataRow := m.flexBox.NewRow()

	// 	if dataRow == nil {
	// 		panic("could not find the table row")
	// 	}
	// 	// Add first padding tagCellbefore adding content cells
	// 	dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	// 	// Ensures addition or deletion truncation from float to int ends up at length 1.0
	// 	if len(row) == 1 {
	// 		row[0].widthPerUnit = 1.0
	// 	}

	// 	// Add content cells
	// 	for j, cell := range row {

	// 		style := m.getCellTextStyle(cell, styleNormal)
	// 		dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style))
	// 		tagCell := dataRow.GetCell(j + 1).SetContent("")
	// 		if tagCell == nil {
	// 			panic("could not find the table cell")
	// 		}
	// 	}

	// 	// Add closing padding cell
	// 	dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	// 	rows = append(rows, dataRow)
	// }

	/* ---------------------------------------------------------------------------*/
	/* ---------------------------------------------------------------------------*/

	// Add closing padding row
	lastRow := m.flexBox.NewRow()
	lastRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG).SetContent(footerText)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	// Add text input interface
	if m.updateType == textInput {
		lastRow.GetCell(1).SetStyle(styleTextInput)
	}

	//Append whatever was added, data binding or textInput/nothing
	rows = append(rows, lastRow)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("A: Insert row below\t\tZ: Delete current row\t\tX: Delete current cell"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("S: Insert cell to the left\tD: Insert cell to the right"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("When you're done building the tag, press B to start binding the data from your CSV file to each cell."))
	rows = append(rows, helpRow)

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling tagBuilderView, we always want to overwrite to refresh the view.

	m.flexBox.SetRows(rows)

	// Highlight the current content row and tagCellas selected
	if (m.tagRowCursor >= 0 && m.tagRowCursor < len(m.tag.tagTable)) &&
		(m.tagCellCursor >= 0 && m.tagCellCursor < len(m.tag.tagTable[m.tagRowCursor])) {

		cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]
		style := m.getCellTextStyle(cell, styleSelected)
		rows[m.fbTagRowCursor()].GetCell(m.fbTagCellCursor()).SetStyle(style)
	}
}

func (m *model) dataBinderView(footerText string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBindList).SetContent("Table Data Binder")).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	rows = append(rows, firstRow)

	/* ---------------------------------------------------------------------------*/
	/* ---------------------------------------------------------------------------*/

	// Add tag rows
	for _, row := range m.tag.tagTable {

		dataRow := m.flexBox.NewRow()

		if dataRow == nil {
			panic("could not find the table row")
		}
		// Add first padding tagCellbefore adding content cells
		dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

		// Ensures addition or deletion truncation from float to int ends up at length 1.0
		if len(row) == 1 {
			row[0].widthPerUnit = 1.0
		}

		// Add content cells
		for _, cell := range row {

			style := m.getCellTextStyle(cell, styleNormal)
			if len(cell.refHeader) > 0 {
				if cell.isFieldName {
					dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style).
						SetContent(cell.refHeader))
				} else {
					dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style).
						SetContent(m.csvData.rows[m.currentTag][cell.refHeader]))
				}

			} else {
				dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style))
			}
		}

		// Add closing padding cell
		dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, dataRow)
	}

	/* ---------------------------------------------------------------------------*/
	/* ---------------------------------------------------------------------------*/

	// Add closing padding row
	lastRow := m.flexBox.NewRow()
	lastRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBindList).SetContent(footerText)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	//Append whatever was added, data binding or textInput/nothing
	rows = append(rows, lastRow)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Space: Bind table data to tag cell\t\t⌫: Skip current binding\t\tEsc: Return to Tag Builder"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("When you're done binding the data, press V to go to Tag Viewer and Printing."))
	rows = append(rows, helpRow)

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling tagBuilderView, we always want to overwrite to refresh the view.

	m.flexBox.SetRows(rows)

	// Highlight the current content row and tagCellas selected
	if (m.tagRowCursor >= 0 && m.tagRowCursor < len(m.tag.tagTable)) &&
		(m.tagCellCursor >= 0 && m.tagCellCursor < len(m.tag.tagTable[m.tagRowCursor])) {

		cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]
		style := m.getCellTextStyle(cell, styleSelected)
		rows[m.fbTagRowCursor()].GetCell(m.fbTagCellCursor()).SetStyle(style)
	}
}

func (m *model) tagViewerView(footerText string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBindList).SetContent("Tag Viewer")).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	rows = append(rows, firstRow)

	/* ---------------------------------------------------------------------------*/
	/* ---------------------------------------------------------------------------*/

	// Add tag rows
	for _, row := range m.tag.tagTable {

		dataRow := m.flexBox.NewRow()

		if dataRow == nil {
			panic("could not find the table row")
		}
		// Add first padding tagCellbefore adding content cells
		dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

		// Ensures addition or deletion truncation from float to int ends up at length 1.0
		if len(row) == 1 {
			row[0].widthPerUnit = 1.0
		}

		// Add content cells
		for _, cell := range row {
			style := m.getCellTextStyle(cell, styleNormal)
			if len(cell.refHeader) > 0 {
				if cell.isFieldName {
					dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style).
						SetContent(cell.refHeader))
				} else {
					dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style).
						SetContent(m.csvData.rows[m.currentTag][cell.refHeader]))
				}
			} else {
				dataRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style))
			}
		}

		// Add closing padding cell
		dataRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, dataRow)
	}

	/* ---------------------------------------------------------------------------*/
	/* ---------------------------------------------------------------------------*/

	// Add closing padding row
	lastRow := m.flexBox.NewRow()
	lastRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBindList).SetContent(footerText)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	// Add text input interface
	if m.updateType == textInput {
		lastRow.GetCell(1).SetStyle(styleTextInput)
	}

	//Append whatever was added, data binding or textInput/nothing
	rows = append(rows, lastRow)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("B: Set/Unset Bold\t\tI: Set/Unset Italic\t\tC: Toggle horizontal center align"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("W: change current cell width (UNSAFE: make sure all cell widths in the row add up to 1.0)"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("K: Previous tag\t\tL: Next tag\t\tP: Generate PDF"))
	rows = append(rows, helpRow)

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling tagBuilderView, we always want to overwrite to refresh the view.

	m.flexBox.SetRows(rows)

	// Highlight the current content row and tagCellas selected
	if (m.tagRowCursor >= 0 && m.tagRowCursor < len(m.tag.tagTable)) &&
		(m.tagCellCursor >= 0 && m.tagCellCursor < len(m.tag.tagTable[m.tagRowCursor])) {

		cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]
		style := m.getCellTextStyle(cell, styleSelected)
		rows[m.fbTagRowCursor()].GetCell(m.fbTagCellCursor()).SetStyle(style)
	}
}
