package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
)

type currentView = int
type caller = int

const (
	WelcomeView currentView = iota
	TagView
	BuildView
	FileLoaderView
	PrintView
)

const (
	leftInsert caller = iota
	rightInsert
)

type Cell struct {
	widthPerUnit float64
	text         string
	centered     bool
	textStyle    string
}
type TagRow []Cell     // each element is a cell
type TagTable []TagRow // each element is a row with cells

type Tag struct {
	// width  float64
	// height float64
	table TagTable
}

type CSVData struct {
	headings []string
	headers  []string
	data     [][]string
}

type model struct {
	currentView    currentView
	withTextInput  bool
	tagCursorRow   int
	tagCursorCell  int
	inputTextValue string
	inputCaller    caller
	textInput      textinput.Model
	flexBox        *flexbox.FlexBox
	tag            Tag
}

func InitialModel() *model {

	dm := model{}
	dm.flexBox = flexbox.New(800, 600)
	dm.tag = Tag{
		// width:  80.0,
		// height: 40.0,
		table: TagTable{
			{
				{widthPerUnit: 1.0,
					text:      "Project",
					centered:  true,
					textStyle: "B"},
			},
			{
				{widthPerUnit: 1.0,
					text:      "Project Decription",
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
					text:      "Field 1 Data",
					centered:  false,
					textStyle: ""},
			},
			{
				{widthPerUnit: 0.3,
					text:      "Name",
					centered:  false,
					textStyle: "B"},
				{widthPerUnit: 0.2,
					text:      "MNI",
					centered:  false,
					textStyle: ""},
				{widthPerUnit: 0.3,
					text:      "Date",
					centered:  false,
					textStyle: "BI"},
				{widthPerUnit: 0.2,
					text:      "2024-12-05",
					centered:  false,
					textStyle: "I"},
			},
		},
	}

	// Cursors start at -1 to avoid starting with a cell selected
	dm.tagCursorRow = 0
	dm.tagCursorCell = 0
	dm.currentView = WelcomeView

	dm.TagView("")

	return &dm
}
