package main

import (
	"slices"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
)

// tagCursorDown move tagTable cursor down (through rows)
func (m *model) tagCursorDown() {

	// m.setCursor()

	if m.tagRowCursor+1 < len(m.tag.tagTable) {

		nextRowCursor := m.tagRowCursor + 1
		nextRowLen := len(m.tag.tagTable[nextRowCursor])

		var nextCellCursor int
		if m.tagCellCursor >= nextRowLen {
			nextCellCursor = nextRowLen - 1
		} else {
			nextCellCursor = m.tagCellCursor
		}

		// m.flexBox.GetRow(m.fbTagRowCursor()).GetCell(m.fbTagCellCursor()).SetStyle(styleNormal)
		// m.flexBox.GetRow(m.fbTagRowCursorFromRef(nextRowCursor)).GetCell(m.fbTagCellCursorFromRef(nextCellCursor)).SetStyle(styleSelected)

		m.tagRowCursor = nextRowCursor
		m.tagCellCursor = nextCellCursor
	}
}

// tagCursorUp move tagTable cursor up (through rows)
func (m *model) tagCursorUp() {

	// m.setCursor()

	if m.tagRowCursor > 0 { // 1 because of padding row

		nextRowCursor := m.tagRowCursor - 1
		nextRowLen := len(m.tag.tagTable[nextRowCursor])

		var nextCellCursor int

		if m.tagCellCursor >= nextRowLen {
			nextCellCursor = nextRowLen - 1
		} else {
			nextCellCursor = m.tagCellCursor
		}

		// m.flexBox.GetRow(m.fbTagRowCursor()).GetCell(m.fbTagCellCursor()).SetStyle(styleNormal)
		// m.flexBox.GetRow(m.fbTagRowCursorFromRef(nextRowCursor)).GetCell(m.fbTagCellCursorFromRef(nextCellCursor)).SetStyle(styleSelected)

		m.tagRowCursor = nextRowCursor
		m.tagCellCursor = nextCellCursor
	}
}

// tagCursorRight move tagTable cursor right (through cells)
func (m *model) tagCursorRight() {

	// m.setCursor()

	if m.tagCellCursor+1 < len(m.tag.tagTable[m.tagRowCursor]) {

		nextCellCursor := m.tagCellCursor + 1

		// m.flexBox.GetRow(m.fbTagRowCursor()).GetCell(m.fbTagCellCursor()).SetStyle(styleNormal)
		// m.flexBox.GetRow(m.fbTagRowCursor()).GetCell(m.fbTagCellCursorFromRef(nextCellCursor)).SetStyle(styleSelected)

		m.tagCellCursor = nextCellCursor
	}
}

// tagCursorLeft move tagTable cursor left (through cells)
func (m *model) tagCursorLeft() {

	// m.setCursor()

	if m.tagCellCursor > 0 {

		nextCellCursor := m.tagCellCursor - 1

		// m.flexBox.GetRow(m.fbTagRowCursor()).GetCell(m.fbTagCellCursor()).SetStyle(styleNormal)
		// m.flexBox.GetRow(m.fbTagRowCursor()).GetCell(m.fbTagCellCursorFromRef(nextCellCursor)).SetStyle(styleSelected)

		m.tagCellCursor = nextCellCursor
	}
}

// AddTagRow adds a row below current row sets a tag field to the csv header
func (m *model) insertTagRow() {

	if len(m.tag.tagTable) == 0 {
		m.tag.tagTable = append(m.tag.tagTable, tagRow{
			{
				widthPerUnit: 1.0,
				refHeader:    "",
				isFieldName:  false,
				centered:     false,
				textStyle:    "",
			}})
	} else {

		nextRowCursor := m.tagRowCursor + 1

		m.tag.tagTable = slices.Insert(m.tag.tagTable, nextRowCursor, tagRow{})

		m.tag.tagTable[nextRowCursor] = append(m.tag.tagTable[nextRowCursor],
			cell{
				widthPerUnit: 1.0,
				refHeader:    "",
				isFieldName:  false,
				centered:     false,
				textStyle:    "",
			})

	}

}

func (m *model) deleteTagRow() {

	// delete current row only if there are elements in the tagTable to avoid
	// panic because of empty tagTable
	if len(m.tag.tagTable) > 1 {

		nextRowCursor := m.tagRowCursor + 1
		m.tag.tagTable = slices.Delete(m.tag.tagTable, m.tagRowCursor, nextRowCursor)

		if m.tagRowCursor >= len(m.tag.tagTable) {
			m.tagRowCursor = len(m.tag.tagTable) - 1
		}

		if m.tagCellCursor >= len(m.tag.tagTable[m.tagRowCursor]) {
			m.tagCellCursor = len(m.tag.tagTable[m.tagRowCursor]) - 1
		}
	}

}

// Add a cell
func (m *model) insertTagCellLeft() {

	// m.setCursor() // sanity

	widthPU := m.inputValues.floatVal
	// widthPU, _ := strconv.ParseFloat(m.getCellWidthValue(), 64)

	if widthPU != 0.0 {

		tempCell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]

		if widthPU != tempCell.widthPerUnit {

			tempCell.widthPerUnit = tempCell.widthPerUnit - widthPU
			m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = tempCell

			// for i, cell := range m.tag.tagTable[m.tagRowCursor] {
			// 	m.tag.tagTable[m.tagRowCursor][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
			// }
			m.tag.tagTable[m.tagRowCursor] = slices.Insert(m.tag.tagTable[m.tagRowCursor], m.tagCellCursor, cell{
				widthPerUnit: widthPU,
				refHeader:    "",
				isFieldName:  false,
				centered:     false,
				textStyle:    "",
			})
		}
	}
}

func (m *model) insertTagCellRight() {

	widthPU := m.inputValues.floatVal

	if widthPU != 0.0 {

		tempCell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]

		if widthPU != tempCell.widthPerUnit {

			tempCell.widthPerUnit = tempCell.widthPerUnit - widthPU
			m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = tempCell

			// for i, cell := range m.tag.tagTable[m.tagRowCursor] {
			// 	m.tag.tagTable[m.tagRowCursor][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
			// }
			nextCellCursor := m.tagCellCursor + 1
			m.tag.tagTable[m.tagRowCursor] = slices.Insert(m.tag.tagTable[m.tagRowCursor], nextCellCursor, cell{
				widthPerUnit: widthPU,
				refHeader:    "",
				isFieldName:  false,
				centered:     false,
				textStyle:    "",
			})
		}

	}
}

func (m *model) deleteTagCell() {

	var row tagRow

	if lenRow := len(m.tag.tagTable[m.tagRowCursor]); lenRow == 1 {
		m.deleteTagRow()
		return
	} else if m.tagCellCursor == lenRow-1 {
		// Just remove the end element
		row = m.tag.tagTable[m.tagRowCursor][:lenRow-1]
		m.tagCellCursor = len(row) - 1
	} else {
		nextCellCursor := m.tagCellCursor + 1

		row = m.tag.tagTable[m.tagRowCursor]
		row = slices.Delete(row, m.tagCellCursor, nextCellCursor)
	}

	m.tag.tagTable[m.tagRowCursor] = row

	lenResizedRow := 0.0
	for _, cell := range m.tag.tagTable[m.tagRowCursor] {
		lenResizedRow += cell.widthPerUnit
	}

	for i, cell := range m.tag.tagTable[m.tagRowCursor] {
		m.tag.tagTable[m.tagRowCursor][i].widthPerUnit = cell.widthPerUnit / float64(lenResizedRow)
	}

}

// Creates textinput instance
func (m *model) setUserInput(caller func()) {
	m.activeInput = true
	m.inputCaller = caller
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 50
	m.textInput = ti
}

func (m *model) unsetUserInput() {
	m.activeInput = false
	m.inputCaller = nil
}

// Flexbox cursor is always calculated from the tagTable, it cannot go out of
// bound because the flexbox is always padded and bigger than the data tagTable.
func (m *model) fbTagRowCursor() int {

	return m.tagRowCursor + 1
}

func (m *model) fbTagCellCursor() int {

	return m.tagCellCursor + 1
}

func (m *model) dataBindToCell() {

	if m.prevCSVHeaderIdx+1 == m.currentCSVHeaderIdx {

		cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]
		cell.refHeader = m.csvData.headers[m.currentCSVHeaderIdx]
		cell.isFieldName = true
		m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = cell
		m.csvData.boundHeaders[m.currentCSVHeaderIdx] = true
		m.prevCSVHeaderIdx++

	} else if m.prevCSVHeaderIdx == m.currentCSVHeaderIdx {
		cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]
		cell.refHeader = m.csvData.headers[m.currentCSVHeaderIdx]
		cell.isFieldName = false
		m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = cell

		m.csvData.boundRows[m.currentCSVHeaderIdx] = true
		m.currentCSVHeaderIdx++
	}

	if m.currentCSVHeaderIdx == len(m.csvData.boundHeaders) {
		m.currentCSVHeaderIdx = 0
		m.prevCSVHeaderIdx = -1
	}

}

func (m *model) skipBindToCell() {
	length := len(m.csvData.boundHeaders)
	if m.currentCSVHeaderIdx < length {
		m.csvData.boundHeaders[m.currentCSVHeaderIdx] = true
		m.currentCSVHeaderIdx++
	} else if m.currentCSVHeaderIdx >= length && m.currentCSVHeaderIdx < length*2 {
		idx := m.currentCSVHeaderIdx - length

		m.csvData.boundRows[idx] = true
		m.currentCSVHeaderIdx++
	}

}

func (m *model) nextTag() {

	nextTag := m.currentTag + 1

	if nextTag < len(m.csvData.rows) {
		m.currentTag = nextTag
	} else {
		m.currentTag = 0 // loop back
	}
}

func (m *model) previousTag() {

	nextTag := m.currentTag - 1

	if nextTag > 0 {
		m.currentTag = nextTag
	} else {
		m.currentTag = len(m.csvData.rows) - 1 // loop forward
	}
}

func (m *model) printCursorDown() {

	if m.printRowCursor+1 <= len(printKeysStatic)+len(printKeysInputs) {

		nextRowCursor := m.printRowCursor + 1

		m.printRowCursor = nextRowCursor
		m.printCellCursor = 2 // always return to first cell
	}
}

//Special case cursor for print screen

func (m *model) printCursorUp() {

	if m.printRowCursor-1 > 0 { // Off from 0 because of padding row

		nextRowCursor := m.printRowCursor - 1

		m.printRowCursor = nextRowCursor
		m.printCellCursor = 2 // always return to first cell

	}
}

// printCursorRight moves cursor right (through cells), visually
func (m *model) printCursorRight() {

	if m.printCellCursor+1 < m.flexBox.GetRow(m.printRowCursor).CellsLen()-1 {

		nextCellCursor := m.printCellCursor + 1
		m.printCellCursor = nextCellCursor
	}
}

// printCursorLeft move cursor left (through cells), visually
func (m *model) printCursorLeft() {

	if m.printCellCursor-1 >= 2 {

		nextCellCursor := m.printCellCursor - 1
		m.printCellCursor = nextCellCursor
	}
}

func (m *model) getPaperSize() string {

	return m.flexBox.GetRow(m.printRowCursor).GetCell(m.printCellCursor).GetContent()
}

func (m *model) resetViewState() {
	m.activeInput = false
	m.currentCSVHeaderIdx = 0
	m.tagRowCursor = 0
	m.tagCellCursor = 0
	m.inputValues = inputValues{}
}

func (m *model) changeCellWidth() {

	floatVal := m.inputValues.floatVal

	if floatVal <= 0.1 && floatVal > 0.8 {
		return
	}

	m.tag.tagTable[m.tagRowCursor][m.tagCellCursor].widthPerUnit = floatVal
}

func (m *model) saveInputValue() {

	strInput := m.textInput.Value()
	floatInput, _ := strconv.ParseFloat(strInput, 64)

	if floatInput > 0.0 {
		m.inputValues.floatVal = floatInput
		m.inputValues.stringVal = strInput
	} else {
		m.inputValues.floatVal = 0.0
		m.inputValues.stringVal = ""
	}
}

func (m *model) saveToCurrentPVSelected() {

	stringVal := m.inputValues.stringVal

	// Both need to be saved otherwise comparison in view is impossible
	m.pVContents.selectedValues[printViewRows[m.printRowCursor]] = stringVal
	m.pVContents.rows[printViewRows[m.printRowCursor]][1] = stringVal

}
