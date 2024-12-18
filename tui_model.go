package main

import (
	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
)

// Definitions

type currentView = int

const (
	welcome1View currentView = iota
	welcome2View
	tagBuilderView
	dataBinderView
	tagViewerView
	printToPDFView
)

// Declarations

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
	// width    float64 // real width in mm
	// height   float64 // real height in mm
	tagTable tagTable
}

type csvData struct {
	headers      []string            // index parity with corresponding bound
	rows         []map[string]string // index parity with corresponding bound
	boundHeaders []bool
	boundRows    []bool
}

type printViewContents struct {
	rows           map[string][]string
	selectedValues map[string]string
}

var paperSizes []string

// var fontTypes []string
// var printKeysRowsStaticSI map[string]int
// var printKeysRowsInputsSI map[string]int
var printKeysStatic map[int]string
var printKeysInputs map[int]string
var printViewRows map[int]string

type inputValues struct {
	floatVal  float64
	stringVal string
}

type model struct {
	currentView         currentView
	inputCaller         func()
	activeInput         bool
	tagRowCursor        int
	tagCellCursor       int
	printRowCursor      int
	printCellCursor     int
	currentTag          int
	currentCSVHeaderIdx int
	prevCSVHeaderIdx    int
	inputValues         inputValues
	pVContents          printViewContents
	textInput           textinput.Model
	tag                 tagRepr
	csvData             csvData
	flexBox             *flexbox.FlexBox
}

func InitialModel() *model {

	dm := model{}
	dm.readCSVFile()
	dm.flexBox = flexbox.New(0, 0)

	dm.tag = tagRepr{
		tagTable: tagTable{},
	}

	dm.inputValues = inputValues{}

	dm.tagRowCursor = 0
	dm.tagCellCursor = 0
	dm.printRowCursor = 1
	dm.printCellCursor = 2
	dm.currentTag = 0
	dm.currentView = welcome1View
	dm.activeInput = false
	dm.currentCSVHeaderIdx = 0
	dm.prevCSVHeaderIdx = -1

	printKeysStatic = map[int]string{1: "paper"}
	// printKeysRowsStatic = map[int]string{1: "paper", 2:"font"}

	printKeysInputs = map[int]string{2: "width", 3: "height", 4: "fontSize"}

	printViewRows = map[int]string{1: "paper", 2: "width", 3: "height", 4: "fontSize"}

	paperSizes = []string{"A4", "Letter"}

	dm.pVContents.rows = map[string][]string{
		"paper": {"Select paper size:"},
		// "fontType": {"Select font type:", ""},
		"width":    {"Enter tag width [mm]:", ""},
		"height":   {"Enter tag height [mm]:", ""},
		"fontSize": {"Enter font size [pt]:", ""},
	}

	dm.pVContents.selectedValues = map[string]string{
		"paper":    "",
		"width":    "",
		"height":   "",
		"fontSize": "",
	}

	dm.pVContents.rows["paper"] = append(dm.pVContents.rows["paper"], paperSizes...)

	dm.csvData.boundRows = make([]bool, len(dm.csvData.headers))
	dm.csvData.boundHeaders = make([]bool, len(dm.csvData.headers))

	dm.tagBuilderView("")

	return &dm
}
