package main

import "slices"

// CursorDown move table cursor down (through rows)
func (m *model) CursorDown() {
	if m.currCursorRow+1 <= len(m.tag.table) {

		m.nextCursorRow = m.currCursorRow + 1

		if m.currCursorCell > len(m.tag.table[m.nextCursorRow-1]) {
			m.nextCursorCell = len(m.tag.table[m.nextCursorRow-1])
		} else {
			m.nextCursorCell = m.currCursorCell
		}

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(m.nextCursorRow).GetCell(m.nextCursorCell).SetStyle(styleSelected)

		m.currCursorCell = m.nextCursorCell
		m.currCursorRow = m.nextCursorRow
	}
}

// CursorUp move table cursor up (through rows)
func (m *model) CursorUp() {
	if m.currCursorRow > 1 { // 1 because of padding row

		m.nextCursorRow = m.currCursorRow - 1
		if m.currCursorCell >= len(m.tag.table[m.nextCursorRow-1]) {
			m.nextCursorCell = len(m.tag.table[m.nextCursorRow-1])
		} else {
			m.nextCursorCell = m.currCursorCell
		}

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(m.nextCursorRow).GetCell(m.nextCursorCell).SetStyle(styleSelected)

		m.currCursorCell = m.nextCursorCell
		m.currCursorRow = m.nextCursorRow
	}
}

// CursorRight move table cursor right (through cells)
func (m *model) CursorRight() {
	if m.currCursorCell+1 <= len(m.tag.table[m.currCursorRow-1]) {

		m.nextCursorCell = m.currCursorCell + 1

		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(m.nextCursorRow).GetCell(m.nextCursorCell).SetStyle(styleSelected)

		m.currCursorCell = m.nextCursorCell
		m.currCursorRow = m.nextCursorRow
	}
}

// CursorLeft move table cursor left (through cells)
func (m *model) CursorLeft() {
	if m.currCursorCell > 1 { // -1 because of padding cell

		m.nextCursorCell = m.currCursorCell - 1
		m.flexBox.GetRow(m.currCursorRow).GetCell(m.currCursorCell).SetStyle(styleNormal)
		m.flexBox.GetRow(m.nextCursorRow).GetCell(m.nextCursorCell).SetStyle(styleSelected)

		m.currCursorCell = m.nextCursorCell
		m.currCursorRow = m.nextCursorRow
	}
}

// AddTagRow adds a row below current row sets a tag field to the csv header
func (m *model) InsertTagRow() {

	m.tag.table = slices.Insert(m.tag.table, m.currCursorRow, TagRow{})
	m.tag.table[m.currCursorRow] = append(
		m.tag.table[m.currCursorRow],
		Cell{
			widthPerUnit: 1.0,
			text:         "AR",
			centered:     false,
			textStyle:    "",
		})
	// m.tag.printStructure()
	m.createRows()
}

// Add a cell, other cells will reduce their size proportional tu their current
// ratio to accomodate the new cell
func (m *model) InsertTagCellLeft(widthPU float64) *model {

	for i, cell := range m.tag.table[m.currCursorRow-1] {
		m.tag.table[m.currCursorRow-1][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
	}
	m.tag.table[m.currCursorRow-1] = slices.Insert(m.tag.table[m.currCursorRow-1], m.currCursorCell-1, Cell{
		widthPerUnit: widthPU,
		text:         "ACL",
		centered:     false,
		textStyle:    "",
	})
	m.createRows()
	return m
}

func (m *model) InsertTagCellRight(widthPU float64) *model {

	for i, cell := range m.tag.table[m.currCursorRow-1] {
		m.tag.table[m.currCursorRow-1][i].widthPerUnit = cell.widthPerUnit - cell.widthPerUnit*widthPU
	}
	m.tag.table[m.currCursorRow-1] = slices.Insert(m.tag.table[m.currCursorRow], m.currCursorCell, Cell{
		widthPerUnit: widthPU,
		text:         "ACR",
		centered:     false,
		textStyle:    "",
	})
	m.createRows()
	return m
}

// // BindData sets a tag field to a CSV column data and skips the header
