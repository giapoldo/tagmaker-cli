package main

import (
	"slices"

	"github.com/charmbracelet/bubbles/textinput"
)

// CursorDown move table cursor down (through rows)
func (m *model) CursorDown() {
	if m.currCursorRow+1 <= len(m.tag.table) {

		nextCursorRow := m.currCursorRow + 1
		var nextCursorCell int

		m.setCursor()

		if m.currCursorCell > len(m.tag.table[m.offsetTableCursorR(1)]) {
			nextCursorCell = len(m.tag.table[m.offsetTableCursorR(1)])
		} else {
			nextCursorCell = m.currCursorCell
		}

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(nextCursorRow).GetCell(nextCursorCell).SetStyle(styleSelected)

		m.currCursorRow = nextCursorRow
		m.currCursorCell = nextCursorCell
	}
}

// CursorUp move table cursor up (through rows)
func (m *model) CursorUp() {
	if m.currCursorRow > 1 { // 1 because of padding row

		nextCursorRow := m.currCursorRow - 1
		var nextCursorCell int

		m.setCursor()

		if m.currCursorCell > len(m.tag.table[m.offsetTableCursorR(-1)]) {
			nextCursorCell = len(m.tag.table[m.offsetTableCursorR(-1)])
		} else {
			nextCursorCell = m.currCursorCell
		}

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(nextCursorRow).GetCell(nextCursorCell).SetStyle(styleSelected)

		m.currCursorRow = nextCursorRow
		m.currCursorCell = nextCursorCell
	}
}

// CursorRight move table cursor right (through cells)
func (m *model) CursorRight() {
	if m.currCursorCell+1 <= len(m.tag.table[m.offsetTableCursorR(0)]) {

		m.setCursor()
		nextCursorCell := m.currCursorCell + 1

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(m.currCursorRow).GetCell(nextCursorCell).SetStyle(styleSelected)

		m.currCursorCell = nextCursorCell
	}
}

// CursorLeft move table cursor left (through cells)
func (m *model) CursorLeft() {
	if m.currCursorCell > 1 { // -1 because of padding cell

		m.setCursor()
		nextCursorCell := m.currCursorCell - 1

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(m.currCursorRow).GetCell(nextCursorCell).SetStyle(styleSelected)

		m.currCursorCell = nextCursorCell
	}
}

// AddTagRow adds a row below current row sets a tag field to the csv header
func (m *model) InsertTagRow() {

	m.tag.table = slices.Insert(m.tag.table, m.offsetTableCursorR(1), TagRow{})
	m.tag.table[m.offsetTableCursorR(1)] = append(
		m.tag.table[m.offsetTableCursorR(1)],
		Cell{
			widthPerUnit: 1.0,
			text:         "AR",
			centered:     false,
			textStyle:    "",
		})
	// m.tag.printStructure()
	m.createRows()
}

func (m *model) DeleteTagRow() {

	// delete current row by skipping it in the append
	m.tag.table = slices.Delete(m.tag.table, m.offsetTableCursorR(0), m.offsetTableCursorR(1))
	m.createRows()
}

// Add a cell, other cells will reduce their size proportional tu their current
// ratio to accomodate the new cell
func (m *model) InsertTagCellLeft(widthPU float64) *model {

	m.setCursor() // sanity

	for i, cell := range m.tag.table[m.offsetTableCursorR(0)] {
		m.tag.table[m.offsetTableCursorR(0)][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
	}
	m.tag.table[m.offsetTableCursorR(0)] = slices.Insert(m.tag.table[m.offsetTableCursorR(0)], m.offsetTableCursorC(0), Cell{
		widthPerUnit: widthPU,
		text:         "",
		centered:     false,
		textStyle:    "",
	})
	m.createRows()
	return m
}

func (m *model) InsertTagCellRight(widthPU float64) *model {

	m.setCursor() // sanity
	m.userInput()

	for i, cell := range m.tag.table[m.offsetTableCursorR(0)] {
		m.tag.table[m.offsetTableCursorR(0)][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
	}
	m.tag.table[m.offsetTableCursorR(0)] = slices.Insert(m.tag.table[m.offsetTableCursorR(0)], m.offsetTableCursorC(0), Cell{
		widthPerUnit: widthPU,
		text:         "",
		centered:     false,
		textStyle:    "",
	})
	m.createRows()
	return m
}

func (m *model) DeleteTagCell() {

	row := m.tag.table[m.offsetTableCursorR(0)]
	widthPU := row[m.offsetTableCursorC(0)].widthPerUnit
	for i, cell := range m.tag.table[m.offsetTableCursorR(0)] {
		m.tag.table[m.offsetTableCursorR(0)][i].widthPerUnit = cell.widthPerUnit + cell.widthPerUnit*widthPU
	}

	row = slices.Delete(row, m.offsetTableCursorC(0), m.offsetTableCursorC(1))

	m.tag.table[m.offsetTableCursorR(0)] = row
	m.createRows()
}

func (m *model) setCursor() {
	if m.currCursorCell == 0 {
		m.currCursorCell = 1
	}

	if m.currCursorRow == 0 {
		m.currCursorRow = 1
	}
	m.createRows()
}

// Offset forwards or backwards from present table cursor (-1 off from currCursor)
func (m *model) offsetTableCursorR(offset int) int {
	return (m.currCursorRow - 1) + offset
}

func (m *model) offsetTableCursorC(offset int) int {
	return (m.currCursorCell - 1) + offset
}

func (m *model) userInput() string {

	ti := textinput.New()
	ti.Placeholder = "Enter width %"
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 20

	m.textInput = ti
	return ""
}
