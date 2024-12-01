package main

import (
	"github.com/76creates/stickers/flexbox"
)

type model struct {
	flexBox        *flexbox.FlexBox
	tag            Tag
	currCursorRow  int
	currCursorCell int
	nextCursorRow  int
	nextCursorCell int
}

func (m *model) createRows() {

	m.flexBox = flexbox.New(0, 0)

	rows := []*flexbox.Row{}
	rows = append(rows, m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG)))

	for _, row := range m.tag.table {
		_fbRow := m.flexBox.NewRow()
		// rows = append(rows, m.flexBox.NewRow())
		// m.flexBox.AddRows(rows)

		// _fbRow := m.flexBox.GetRow(i + 1) // +1 because of row padding

		if _fbRow == nil {
			panic("could not find the table row")
		}

		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

		for j, cell := range row {
			_fbRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(styleNormal))
			_cell := _fbRow.GetCell(j + 1).SetContent(cell.text) // +1 because of cell padding
			if _cell == nil {
				panic("could not find the table cell")
			}
		}
		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, _fbRow)
	}
	rows = append(rows, m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 1).SetStyle(styleBG)))

	rows[1].GetCell(1).SetStyle(styleSelected)

	m.flexBox.AddRows(rows)
}

func InitialModel() *model {

	dm := model{}

	dm.tag = Tag{
		width:  80.0,
		height: 40.0,
		table: TagTable{
			{
				{widthPerUnit: 1.0,
					text:  "Title 1",
					title: true,
					style: "B"},
			},
			{
				{widthPerUnit: 1.0,
					text:  "Subtitle",
					title: true,
					style: ""},
			}, {
				{widthPerUnit: 1.0,
					text:  "",
					title: false,
					style: ""},
			},
			{
				{widthPerUnit: 0.5,
					text:  "Field 1",
					title: false,
					style: "B"},
				{widthPerUnit: 0.5,
					text:  "UTF8 Dátå 1",
					title: false,
					style: ""},
			},
			{
				{widthPerUnit: 0.3,
					text:  "Nombre",
					title: false,
					style: "B"},
				{widthPerUnit: 0.2,
					text:  "GMG",
					title: false,
					style: ""},
				{widthPerUnit: 0.3,
					text:  "Fecha",
					title: false,
					style: "B"},
				{widthPerUnit: 0.2,
					text:  "2024",
					title: false,
					style: ""},
			},
		},
	}

	// all 4 cursor states start at 1 because of padding rows and cells.
	dm.currCursorRow = 1
	dm.currCursorCell = 1
	dm.nextCursorRow = 1
	dm.nextCursorCell = 1
	dm.createRows()

	return &dm
}
