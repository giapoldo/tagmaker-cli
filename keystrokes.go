package main

import (
	"slices"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
)

// CursorDown move table cursor down (through rows)
func (m *model) CursorDown() {

	// m.setCursor()

	if m.currCursorRow+1 < len(m.tag.table) {

		nextCursorRow := m.currCursorRow + 1
		nextRowLen := len(m.tag.table[nextCursorRow])

		var nextCursorCell int
		if m.currCursorCell >= nextRowLen {
			nextCursorCell = nextRowLen - 1
		} else {
			nextCursorCell = m.currCursorCell
		}

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRowFromRef(nextCursorRow)).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.currCursorRow = nextCursorRow
		m.currCursorCell = nextCursorCell
	}
}

// CursorUp move table cursor up (through rows)
func (m *model) CursorUp() {

	// m.setCursor()

	if m.currCursorRow > 0 { // 1 because of padding row

		nextCursorRow := m.currCursorRow - 1
		nextRowLen := len(m.tag.table[nextCursorRow])

		var nextCursorCell int

		if m.currCursorCell >= nextRowLen {
			nextCursorCell = nextRowLen - 1
		} else {
			nextCursorCell = m.currCursorCell
		}

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRowFromRef(nextCursorRow)).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.currCursorRow = nextCursorRow
		m.currCursorCell = nextCursorCell
	}
}

// CursorRight move table cursor right (through cells)
func (m *model) CursorRight() {

	// m.setCursor()

	if m.currCursorCell+1 < len(m.tag.table[m.currCursorRow]) {

		nextCursorCell := m.currCursorCell + 1

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.currCursorCell = nextCursorCell
	}
}

// CursorLeft move table cursor left (through cells)
func (m *model) CursorLeft() {

	// m.setCursor()

	if m.currCursorCell > 0 {

		nextCursorCell := m.currCursorCell - 1

		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCell()).SetStyle(styleNormal)
		m.flexBox.GetRow(m.FBCursorRow()).GetCell(m.FBCursorCellFromRef(nextCursorCell)).SetStyle(styleSelected)

		m.currCursorCell = nextCursorCell
	}
}

// AddTagRow adds a row below current row sets a tag field to the csv header
func (m *model) InsertTagRow() {

	nextCursorRow := m.currCursorRow + 1

	m.tag.table = slices.Insert(m.tag.table, nextCursorRow, TagRow{})

	m.tag.table[nextCursorRow] = append(m.tag.table[nextCursorRow],
		Cell{
			widthPerUnit: 1.0,
			text:         "",
			centered:     false,
			textStyle:    "",
		})
	// m.tag.printStructure()
	// m.createRows("")
}

func (m *model) DeleteTagRow() {

	// delete current row only if there are elements in the table to avoid
	// panic because of empty table
	if len(m.tag.table) > 1 {

		nextCursorRow := m.currCursorRow + 1
		m.tag.table = slices.Delete(m.tag.table, m.currCursorRow, nextCursorRow)

		if m.currCursorRow >= len(m.tag.table) {
			m.currCursorRow = len(m.tag.table) - 1
		}

		if m.currCursorCell >= len(m.tag.table[m.currCursorRow]) {
			m.currCursorCell = len(m.tag.table[m.currCursorRow]) - 1
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
		for i, cell := range m.tag.table[m.currCursorRow] {
			m.tag.table[m.currCursorRow][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
		}
		m.tag.table[m.currCursorRow] = slices.Insert(m.tag.table[m.currCursorRow], m.currCursorCell, Cell{
			widthPerUnit: widthPU,
			text:         "",
			centered:     false,
			textStyle:    "",
		})
		// m.createRows("")
	}
}

func (m *model) InsertTagCellRight() {

	// m.setCursor() // sanity
	widthPU := m.getCellInputValue()
	// widthPU, _ := strconv.ParseFloat(m.getCellInputValue(), 64)

	if widthPU != 0.0 {
		for i, cell := range m.tag.table[m.currCursorRow] {
			m.tag.table[m.currCursorRow][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
		}
		nextCursorCell := m.currCursorCell + 1
		m.tag.table[m.currCursorRow] = slices.Insert(m.tag.table[m.currCursorRow], nextCursorCell, Cell{
			widthPerUnit: widthPU,
			text:         "",
			centered:     false,
			textStyle:    "",
		})
	}
}

func (m *model) DeleteTagCell() {

	if len(m.tag.table[m.currCursorRow]) == 1 {
		m.DeleteTagRow()
	} else {

		nextCursorCell := m.currCursorCell + 1

		row := m.tag.table[m.currCursorRow]
		row = slices.Delete(row, m.currCursorCell, nextCursorCell)

		lenResizedRow := 0.0
		for _, cell := range m.tag.table[m.currCursorRow] {
			lenResizedRow += cell.widthPerUnit
		}

		for i, cell := range m.tag.table[m.currCursorRow] {
			m.tag.table[m.currCursorRow][i].widthPerUnit = cell.widthPerUnit / float64(lenResizedRow)
		}

		m.tag.table[m.currCursorRow] = row
		// m.createRows("")
	}
}

// func (m *model) setCursor() {
// 	if m.currCursorCell < 0 {
// 		m.currCursorCell = 0
// 	}

// 	if m.currCursorRow < 0 {
// 		m.currCursorRow = 0
// 	}
// 	// m.createRows("")
// }

// Set if choosing size of cell or binding data? Options arw "cell" and "binding"
func (m *model) SetCellInput(caller string) {
	m.textInputVisibility = true
	m.inputCaller = caller
	ti := textinput.New()
	ti.Placeholder = "Enter width per unit (0.20~0.80) and press Enter"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	m.textInput = ti
}

func (m *model) UnSetUserInput() {
	m.textInputVisibility = false
}

func (m *model) getCellInputValue() float64 {

	inputValue, _ := strconv.ParseFloat(m.textInput.Value(), 64)

	if inputValue > 0.1 && inputValue <= 0.8 {
		return inputValue
	}
	return 0.0
}

// // Offset forwards or backwards from present table cursor (-1 off from currCursor)
// func (m *model) offsetTableCursorR(offset int) int {
// 	return (m.currCursorRow - 1) + offset
// }

// func (m *model) calcTableCursorC(offset int) int {
// 	return (m.currCursorCell - 1) + offset
// }

// Flexbox cursor is always calculated from the table, it cannot go out of
// bound because the flexbox is always padded and bigger than the data table.
func (m *model) FBCursorRow() int {

	return m.currCursorRow + 1
}

func (m *model) FBCursorRowFromRef(ref int) int {

	return ref + 1
}

func (m *model) FBCursorCell() int {

	return m.currCursorCell + 1
}

func (m *model) FBCursorCellFromRef(ref int) int {

	return ref + 1
}
