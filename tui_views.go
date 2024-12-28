package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/lipgloss"
)

func (m *model) appendPaddedRow(rows []*flexbox.Row, content string, contentStyle lipgloss.Style) []*flexbox.Row {

	row := m.flexBox.NewRow()
	row.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(contentStyle).SetContent(content)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	rows = append(rows, row)

	return rows
}

func (m *model) welcome1View() {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(12)

	// Add top padding row
	rows = m.appendPaddedRow(rows, "", styleBG)

	welcomeText := "Welcome to TagBuilder.\n\nThis tool helps you build a tag format for classification of elements in collections, samples or fieldwork findings around your tabulated data.\n\nInitially for archaeological work, but hopefully useful for other stuff too.\n\nCheck the provided CSV table in the \"Input\" folder for the general format.\n\nPress Ctrl+C or Q to quit at any time, or any other key to continue."

	rows = m.appendPaddedRow(rows, welcomeText, styleWelcome)

	// Add bottom padding row
	rows = m.appendPaddedRow(rows, "", styleBG)

	m.flexBox.SetRows(rows)
}

func (m *model) welcome2View() {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(12)

	// Top padding
	rows = m.appendPaddedRow(rows, "", styleBG)

	welcomeText := "Press A to add your first row.\nThen use the help below to construct the rest of the tag and navigate."

	rows = m.appendPaddedRow(rows, welcomeText, styleWelcome)

	// Bottom padding
	rows = m.appendPaddedRow(rows, "", styleBG)

	m.flexBox.SetRows(rows)
}

func (m *model) tagBuilderView(footerText string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	rows = m.appendPaddedRow(rows, "Tag Builder", styleBindList)

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

	// Add text input interface to last row
	if m.activeInput {
		rows = m.appendPaddedRow(rows, footerText, stylePrintTextInput)
	} else {
		rows = m.appendPaddedRow(rows, footerText, styleBG)
	}

	// //Append whatever was added, data binding or textInput/nothing
	// rows = append(rows, lastRow)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("A: Insert row below\t\tX: Delete selected cell"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("S: Insert cell to the left\tD: Insert cell to the right"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("N: Go to Data Binding screen"))
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

	rows = m.appendPaddedRow(rows, "Table Data Binder", styleBindList)
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

	rows = m.appendPaddedRow(rows, footerText, styleBindList)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Space: Bind table data to tag cell\t\tBackspace: Skip current binding"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("ESC: Return to previous screen"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("N: Go to Tag Viewer and Styler screen"))
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

	rows = m.appendPaddedRow(rows, "Tag Viewer", styleBindList)

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

	// Add bottom row with or without text input interface style
	if m.activeInput {
		rows = m.appendPaddedRow(rows, footerText, styleTextInput)
	} else {
		rows = m.appendPaddedRow(rows, footerText, styleBindList)
	}

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("B: Set/Unset Bold\t\tI: Set/Unset Italic\t\tC: Toggle horizontal center align"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("W: change current cell width (UNSAFE: make sure all cell widths in the row add up to 1.0)"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("K: Previous tag\t\tL: Next tag"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("N: Go to Print Screen\t\tESC: Return to previous screen"))
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

func (m *model) printToPDFView(userInput string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	// Header
	rows = m.appendPaddedRow(rows, "Print to PDF", styleBindList)

	// Body

	// Static text

	pVRow := m.pVContents.rows

	var row *flexbox.Row

	for i := 1; i <= len(printViewRows); i++ { // keys in printViewRows are 1 indexed
		row = m.flexBox.NewRow()
		row.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		lenRow := len(pVRow[printViewRows[i]])

		for j, element := range pVRow[printViewRows[i]] {
			if j == 0 { // first cell
				row.AddCells(flexbox.NewCell(100/lenRow, 1).SetStyle(styleNormal).SetContent(element))
			} else {

				if len(m.pVContents.selectedValues[printViewRows[i]]) > 0 && m.pVContents.selectedValues[printViewRows[i]] == element { // tests if what has been saved is the current element to color it
					row.AddCells(flexbox.NewCell(100/lenRow, 1).SetStyle(stylePermaSelected).SetContent(element))
				} else if m.activeInput && m.printCellCursor == j+1 && m.printRowCursor == i { // if text input is active
					row.AddCells(flexbox.NewCell(100/lenRow, 1).SetStyle(styleTextInput).SetContent(userInput))
				} else {
					row.AddCells(flexbox.NewCell(100/lenRow, 1).SetStyle(styleNormal).SetContent(element))
				}

			}

		}
		row.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, row)
	}

	// Footer
	rows = m.appendPaddedRow(rows, "", styleBG)

	// Help Rows
	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Enter: Activate or set user input\t\tP: Generate PDF"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Esc: Return to previous screen"))
	rows = append(rows, helpRow)

	m.flexBox.SetRows(rows)

	// Highlight the current row and cell as selected
	// rows are 1-indexed (> 0) because of padding
	// cells are 2-indexed (> 1) because of padding and row fieldnames

	if m.printRowCursor > 0 && m.printRowCursor <= (len(printKeysStatic)+len(printKeysInputs)) &&
		m.printCellCursor > 1 && m.printCellCursor < m.flexBox.GetRow(m.printRowCursor).CellsLen()-1 {
		m.flexBox.GetRow(m.printRowCursor).GetCell(m.printCellCursor).SetStyle(styleSelected)
	}
}
