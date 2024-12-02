package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	textInput      textinput.Model
	flexBox        *flexbox.FlexBox
	tag            Tag
	currCursorRow  int
	currCursorCell int
}

func (m *model) createRows( /*setRows bool*/ ) {

	rows := []*flexbox.Row{}

	// Add first padding row before adding tag rows
	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG))
	rows = append(rows, firstRow)

	// Add tag rows
	for _, row := range m.tag.table {
		_fbRow := m.flexBox.NewRow()

		if _fbRow == nil {
			panic("could not find the table row")
		}
		// Add first padding cell before adding content cells
		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		// Add content cells
		for j, cell := range row {
			_fbRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(styleNormal))
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
	lastRow.AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG))
	rows = append(rows, lastRow)

	// Highlight the current content row and cell as selected
	if m.flexBox.RowsLen() > 0 {
		if m.currCursorRow > 0 || m.currCursorRow < m.flexBox.RowsLen()-1 ||
			m.currCursorCell > 0 || m.currCursorCell < m.flexBox.GetRow(m.currCursorRow).CellsLen() {

			rows[m.currCursorRow].GetCell(m.currCursorCell).SetStyle(styleSelected)
		}
	}

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling CreateRows, we always want to overwrite to refresh the view.
	m.flexBox.SetRows(rows)
}

func InitialModel() *model {

	dm := model{}
	dm.flexBox = flexbox.New(0, 0)
	dm.flexBox.LockRowHeight(4)
	dm.tag = Tag{
		// width:  80.0,
		// height: 40.0,
		table: TagTable{
			{
				{widthPerUnit: 1.0,
					text:      "Title 1",
					centered:  true,
					textStyle: "B"},
			},
			{
				{widthPerUnit: 1.0,
					text:      "Subtitle",
					centered:  true,
					textStyle: ""},
			}, {
				{widthPerUnit: 1.0,
					text:      "",
					centered:  false,
					textStyle: ""},
			},
			{
				{widthPerUnit: 0.5,
					text:      "Field 1",
					centered:  false,
					textStyle: "B"},
				{widthPerUnit: 0.5,
					text:      "UTF8 Dátå 1",
					centered:  false,
					textStyle: ""},
			},
			{
				{widthPerUnit: 0.3,
					text:      "Nombre",
					centered:  false,
					textStyle: "B"},
				{widthPerUnit: 0.2,
					text:      "GMG",
					centered:  false,
					textStyle: ""},
				{widthPerUnit: 0.3,
					text:      "Fecha",
					centered:  false,
					textStyle: "B"},
				{widthPerUnit: 0.2,
					text:      "2024",
					centered:  false,
					textStyle: ""},
			},
		},
	}

	// all 4 cursor states start at 1 because of padding rows and cells.
	dm.currCursorRow = 0
	dm.currCursorCell = 0

	dm.createRows()

	return &dm
}
