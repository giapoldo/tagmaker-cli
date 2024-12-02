package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Cell struct {
	widthPerUnit float64
	text         string
	centered     bool
	textStyle    string
}
type TagRow []Cell     // each element is a cell
type TagTable []TagRow // each element is a row with cells

type Tag struct { // cambiar model por Tag si vuelvo a GUI en vez de TUI
	// width  float64 // in units of measurement
	// height float64 // in unirs of measurement
	// lrMargin float64
	// tbMargin float64
	table TagTable // table shape metadata
}

func (t Tag) printStructure() {
	s := ""
	for i, row := range t.table {
		s += fmt.Sprintf("Row %d: ", i)
		for _, cell := range row {
			s += fmt.Sprintf("{widht: %1.1f}, ", cell.widthPerUnit)
		}
		tea.LogToFile("debug.log", s)
		s = ""
	}

}
