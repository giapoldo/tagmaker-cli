package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	textInput           textinput.Model
	textInputVisibility bool
	textValue           string
	inputCaller         string
	flexBox             *flexbox.FlexBox
	tag                 Tag
	currCursorRow       int
	currCursorCell      int
}

type CSVData struct {
	headers []string
	data    [][]string
}

func InitialModel() *model {

	dm := model{}
	dm.flexBox = flexbox.New(800, 600)
	dm.flexBox.LockRowHeight(5)
	dm.tag = Tag{
		// width:  80.0,
		// height: 40.0,
		table: TagTable{
			{
				{widthPerUnit: 1.0,
					text:      "Collection",
					centered:  true,
					textStyle: "B"},
			},
			{
				{widthPerUnit: 1.0,
					text:      "Milestone",
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
					text:      "UTF8 1",
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
					textStyle: "BI"},
				{widthPerUnit: 0.2,
					text:      "2024",
					centered:  false,
					textStyle: "I"},
			},
		},
	}

	// Cursors start at -1 to avoid starting with a cell selected
	dm.currCursorRow = 0
	dm.currCursorCell = 0

	dm.createRows("")

	return &dm
}
