package main

import (
	"slices"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
)

// CursorDown move table cursor down (through rows)
func (m *model) CursorDown() {

	// m.setCursor()

	if m.tagCursorRow+1 < len(m.tag.table) {

		nextCursorRow := m.tagCursorRow + 1
		nextRowLen := len(m.tag.table[nextCursorRow])

		var nextCursorCell int
		if m.tagCursorCell >= nextRowLen {
			nextCursorCell = nextRowLen - 1
		} else {
			nextCursorCell = m.tagCursorCell
		}

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRowFromRef(nextCursorRow)).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.tagCursorRow = nextCursorRow
		m.tagCursorCell = nextCursorCell
	}
}

// CursorUp move table cursor up (through rows)
func (m *model) CursorUp() {

	// m.setCursor()

	if m.tagCursorRow > 0 { // 1 because of padding row

		nextCursorRow := m.tagCursorRow - 1
		nextRowLen := len(m.tag.table[nextCursorRow])

		var nextCursorCell int

		if m.tagCursorCell >= nextRowLen {
			nextCursorCell = nextRowLen - 1
		} else {
			nextCursorCell = m.tagCursorCell
		}

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRowFromRef(nextCursorRow)).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.tagCursorRow = nextCursorRow
		m.tagCursorCell = nextCursorCell
	}
}

// CursorRight move table cursor right (through cells)
func (m *model) CursorRight() {

	// m.setCursor()

	if m.tagCursorCell+1 < len(m.tag.table[m.tagCursorRow]) {

		nextCursorCell := m.tagCursorCell + 1

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.tagCursorCell = nextCursorCell
	}
}

// CursorLeft move table cursor left (through cells)
func (m *model) CursorLeft() {

	// m.setCursor()

	if m.tagCursorCell > 0 {

		nextCursorCell := m.tagCursorCell - 1

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.tagCursorCell = nextCursorCell
	}
}

// AddTagRow adds a row below current row sets a tag field to the csv header
func (m *model) InsertTagRow() {

	nextCursorRow := m.tagCursorRow + 1

	m.tag.table = slices.Insert(m.tag.table, nextCursorRow, TagRow{})

	m.tag.table[nextCursorRow] = append(m.tag.table[nextCursorRow],
		Cell{
			widthPerUnit: 1.0,
			text:         "",
			centered:     false,
			textStyle:    "",
		})
	// m.tag.printStructure()
	// m.TagView("")
}

func (m *model) DeleteTagRow() {

	// delete current row only if there are elements in the table to avoid
	// panic because of empty table
	if len(m.tag.table) > 1 {

		nextCursorRow := m.tagCursorRow + 1
		m.tag.table = slices.Delete(m.tag.table, m.tagCursorRow, nextCursorRow)

		if m.tagCursorRow >= len(m.tag.table) {
			m.tagCursorRow = len(m.tag.table) - 1
		}

		if m.tagCursorCell >= len(m.tag.table[m.tagCursorRow]) {
			m.tagCursorCell = len(m.tag.table[m.tagCursorRow]) - 1
		}
	}

}

// Add a cell, other cells will reduce their size proportional tu their current
// ratio to accomodate the new cell
func (m *model) InsertTagCellLeft() {

	// m.setCursor() // sanity

	widthPU := m.getCellInputValue()
	// widthPU, _ := strconv.ParseFloat(m.getCellInputValue(), 64)

	if widthPU != 0.0 {
		for i, cell := range m.tag.table[m.tagCursorRow] {
			m.tag.table[m.tagCursorRow][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
		}
		m.tag.table[m.tagCursorRow] = slices.Insert(m.tag.table[m.tagCursorRow], m.tagCursorCell, Cell{
			widthPerUnit: widthPU,
			text:         "",
			centered:     false,
			textStyle:    "",
		})
		// m.TagView("")
	}
}

func (m *model) InsertTagCellRight() {

	// m.setCursor() // sanity
	widthPU := m.getCellInputValue()
	// widthPU, _ := strconv.ParseFloat(m.getCellInputValue(), 64)

	if widthPU != 0.0 {
		for i, cell := range m.tag.table[m.tagCursorRow] {
			m.tag.table[m.tagCursorRow][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
		}
		nextCursorCell := m.tagCursorCell + 1
		m.tag.table[m.tagCursorRow] = slices.Insert(m.tag.table[m.tagCursorRow], nextCursorCell, Cell{
			widthPerUnit: widthPU,
			text:         "",
			centered:     false,
			textStyle:    "",
		})
	}
}

func (m *model) DeleteTagCell() {

	var row TagRow

	if lenRow := len(m.tag.table[m.tagCursorRow]); lenRow == 1 {
		m.DeleteTagRow()
		return
	} else if m.tagCursorCell == lenRow-1 {
		// Just remove the end element
		row = m.tag.table[m.tagCursorRow][:lenRow-1]
		m.tagCursorCell = len(row) - 1
	} else {
		nextCursorCell := m.tagCursorCell + 1

		row = m.tag.table[m.tagCursorRow]
		row = slices.Delete(row, m.tagCursorCell, nextCursorCell)
	}
	m.tag.table[m.tagCursorRow] = row

	lenResizedRow := 0.0
	for _, cell := range m.tag.table[m.tagCursorRow] {
		lenResizedRow += cell.widthPerUnit
	}

	for i, cell := range m.tag.table[m.tagCursorRow] {
		m.tag.table[m.tagCursorRow][i].widthPerUnit = cell.widthPerUnit / float64(lenResizedRow)
	}

	// m.TagView("")

}

// Set if choosing size of cell or binding data? Options arw "cell" and "binding"
func (m *model) SetCellInput(callerFunc caller) {
	m.withTextInput = true
	m.inputCaller = callerFunc
	ti := textinput.New()
	ti.Placeholder = "Enter width per unit (0.20~0.80) and press Enter"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	m.textInput = ti
}

func (m *model) UnSetUserInput() {
	m.withTextInput = false
}

func (m *model) getCellInputValue() float64 {

	inputValue, _ := strconv.ParseFloat(m.textInput.Value(), 64)

	if inputValue > 0.1 && inputValue <= 0.8 {
		return inputValue
	}
	return 0.0
}

// Flexbox cursor is always calculated from the table, it cannot go out of
// bound because the flexbox is always padded and bigger than the data table.
func (m *model) FBCursorRow() int {

	return m.tagCursorRow + 1
}

func (m *model) FBCursorRowFromRef(ref int) int {

	return ref + 1
}

func (m *model) FBCursorCell() int {

	return m.tagCursorCell + 1
}

func (m *model) FBCursorCellFromRef(ref int) int {

	return ref + 1
}

func (m *model) BindCSVDataToTag() {

}
