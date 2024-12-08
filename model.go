package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
)

type currentView = int
type caller = int
type tagBuilderActions = int

const (
	welcome1View currentView = iota
	welcome2View
	tagBuilderView
	dataBinderView
	tagViewerView
	printView
)

const (
	cellLeftInsert caller = iota
	cellRightInsert
	rowInsert
	changeCellWidth
)

const (
	textInput tagBuilderActions = iota
	normal
)

type cell struct {
	refHeader    string
	isFieldName  bool // is it a fieldname or it's data?
	widthPerUnit float64
	centered     bool
	textStyle    string
}

type tagRow []cell     // each element is a cell
type tagTable []tagRow // each element is a row with cells

type tagRepr struct {
	// width  float64 // real width in mm
	// height float64 // real height in mm
	// currentTag int
	tagTable tagTable
}

// var rows map[int][]map[string]string

type csvData struct {
	headers      []string            // index parity with corresponding bound
	rows         []map[string]string // index parity with corresponding bound
	boundHeaders []bool
	boundRows    []bool
}

type model struct {
	currentView         currentView
	updateType          tagBuilderActions
	inputCaller         caller
	tagRowCursor        int
	tagCellCursor       int
	currentTag          int
	currentCSVHeaderIdx int
	lastCSVHeaderIdx    int
	inputValue          string
	textInput           textinput.Model
	// tag                []tagRepr
	tag     tagRepr
	csvData csvData
	flexBox *flexbox.FlexBox
}

func InitialModel() *model {

	dm := model{}
	dm.readCSVFile()
	dm.flexBox = flexbox.New(0, 0)
	// dm.tag = make([]tagRepr, len(dm.csvData.rows)) // .data holds the exact amount of tag

	// dm.tag = append(dm.tag, tagRepr{
	// 	tagTable: tagTable{},
	// })

	dm.tag = tagRepr{
		tagTable: tagTable{},
	}

	// Cursors start at -1 to avoid starting with a tagCellselected
	dm.tagRowCursor = 0
	dm.tagCellCursor = 0
	dm.currentTag = 0
	dm.currentView = welcome1View
	dm.updateType = normal
	dm.currentCSVHeaderIdx = 0
	dm.lastCSVHeaderIdx = -1

	dm.csvData.boundRows = make([]bool, len(dm.csvData.headers))
	dm.csvData.boundHeaders = make([]bool, len(dm.csvData.headers))

	dm.tagBuilderView("")

	return &dm
}
