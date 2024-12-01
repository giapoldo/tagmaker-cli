package main

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

// // GetCursorLocation returns the current x,y position of the cursor
// func (r *Table) GetCursorLocation() (int, int) {
// 	return r.cursorIndexX, r.cursorIndexY
// }

// // GetCursorValue returns the string of the cell under the cursor
// func (r *Table) GetCursorValue() string {
// 	// handle 0 rows situation and when table is not active
// 	if len(r.filteredRows) == 0 || r.cursorIndexX < 0 || r.cursorIndexY < 0 {
// 		return ""
// 	}
// 	return getStringFromOrdered(r.filteredRows[r.cursorIndexY][r.cursorIndexX])
// }
