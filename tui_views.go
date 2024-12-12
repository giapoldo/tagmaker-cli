package main

import (
	"strconv"

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

	welcomeText := "Welcome to TagBuilder.\n\nThis tool helps you build a tag format for classification of elements in collections, samples or fieldwork findings around your tabulated data.\n\nInitially for museological work, but use it for whatever yo want.\n\nCheck the provided CSV table in the \"Input\" folder for the format.\n\nPress Ctrl+C or Q to quit or any other key to continue"

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

	welcomeText := "Press A to add your first row.\nthen use the help below to construct the rest of the tag."

	rows = m.appendPaddedRow(rows, welcomeText, styleWelcome)

	// Bottom padding
	rows = m.appendPaddedRow(rows, "", styleBG)

	m.flexBox.SetRows(rows)
}

func (m *model) tagBuilderView(footerText string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	rows = m.appendPaddedRow(rows, "Tag Builder", styleBindList)

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
	rows = m.appendPaddedRow(rows, footerText, styleBG)

	// Add text input interface to last row
	if m.updateType == textInput {
		rows[len(rows)-1].GetCell(1).SetStyle(styleTextInput)
	}

	// //Append whatever was added, data binding or textInput/nothing
	// rows = append(rows, lastRow)

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
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Space: Bind table data to tag cell\t\tBackspace: Skip current binding\t\tEsc: Return to the Tag Builder"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("When you're done binding the data, press V to go to the Tag Viewer."))
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

	// Add bottom row and text input interface if needed
	if m.updateType == textInput {
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

func (m *model) printToPDFView(userInput string) {

	rows := []*flexbox.Row{}
	m.flexBox.LockRowHeight(3)

	pageSizeText := "Select page size:"
	tagWidthText := "Enter tag width (mm):"
	tagHeightText := "Enter tag height (mm):"
	fontSizeText := "Enter font size (pt):"

	rows = m.appendPaddedRow(rows, "Print to PDF", styleBindList)

	pageSizeRow := m.flexBox.NewRow()
	pageSizeRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	pageSizeRow.AddCells(flexbox.NewCell(34, 1).SetStyle(styleNormal).SetContent(pageSizeText))
	if m.paperSize == "A4" {
		pageSizeRow.AddCells(flexbox.NewCell(33, 1).SetStyle(stylePermaSelected).SetContent("A4"))
	} else {
		pageSizeRow.AddCells(flexbox.NewCell(33, 1).SetStyle(styleNormal).SetContent("A4"))
	}
	if m.paperSize == "Letter" {
		pageSizeRow.AddCells(flexbox.NewCell(33, 1).SetStyle(stylePermaSelected).SetContent("Letter"))
	} else {

		pageSizeRow.AddCells(flexbox.NewCell(33, 1).SetStyle(styleNormal).SetContent("Letter"))
	}
	pageSizeRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	tagWidthRow := m.flexBox.NewRow()
	tagWidthRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	tagWidthRow.AddCells(flexbox.NewCell(34, 1).SetStyle(styleNormal).SetContent(tagWidthText))
	if m.tag.width > 0 {
		tagWidthRow.AddCells(flexbox.NewCell(66, 1).SetStyle(stylePermaSelected).SetContent(strconv.FormatFloat(m.tag.width, 'f', -1, 64)))
	} else {
		// Add text input interface
		if m.printRowCursor == 2 {
			tagWidthRow.AddCells(flexbox.NewCell(66, 1).SetStyle(stylePrintTextInput).SetContent(userInput))
		} else {
			tagWidthRow.AddCells(flexbox.NewCell(66, 1).SetStyle(styleNormal).SetContent(""))
		}
	}
	tagWidthRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	tagHeightRow := m.flexBox.NewRow()
	tagHeightRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	tagHeightRow.AddCells(flexbox.NewCell(34, 1).SetStyle(styleNormal).SetContent(tagHeightText))
	if m.tag.height > 0 {
		tagHeightRow.AddCells(flexbox.NewCell(66, 1).SetStyle(stylePermaSelected).SetContent(strconv.FormatFloat(m.tag.height, 'f', -1, 64)))
	} else {
		// Add text input interface
		if m.printRowCursor == 3 {
			tagHeightRow.AddCells(flexbox.NewCell(66, 1).SetStyle(stylePrintTextInput).SetContent(userInput))
		} else {
			tagHeightRow.AddCells(flexbox.NewCell(66, 1).SetStyle(styleNormal).SetContent(""))
		}
	}
	tagHeightRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	fontSizeRow := m.flexBox.NewRow()
	fontSizeRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	fontSizeRow.AddCells(flexbox.NewCell(34, 1).SetStyle(styleNormal).SetContent(fontSizeText))
	if m.tag.fontSize > 0 {
		fontSizeRow.AddCells(flexbox.NewCell(66, 1).SetStyle(stylePermaSelected).SetContent(strconv.FormatFloat(m.tag.fontSize, 'f', -1, 64)))
	} else {
		// Add text input interface
		if m.printRowCursor == 4 {
			fontSizeRow.AddCells(flexbox.NewCell(66, 1).SetStyle(stylePrintTextInput).SetContent(userInput))
		} else {
			fontSizeRow.AddCells(flexbox.NewCell(66, 1).SetStyle(styleNormal).SetContent(""))
		}
	}
	fontSizeRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	rows = append(rows, pageSizeRow, tagWidthRow, tagHeightRow, fontSizeRow)

	// Add closing padding row
	rows = m.appendPaddedRow(rows, "", styleBG)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Arrows to move"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Space: Activate user input\t\tEnter: Set selection or input\t\tP: Generate PDF"))
	rows = append(rows, helpRow)
	helpRow = m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).SetContent("Esc: Return to Tag Viewer"))
	rows = append(rows, helpRow)

	m.flexBox.SetRows(rows)

	// Highlight the current content row and tagCellas selected
	if m.printRowCursor > 0 && m.printRowCursor < m.flexBox.RowsLen() &&
		m.printCellCursor > 0 && m.printCellCursor < m.flexBox.GetRow(m.printRowCursor).CellsLen()-1 {

		m.flexBox.GetRow(m.printRowCursor).GetCell(m.printCellCursor).SetStyle(styleSelected)
	}
}
