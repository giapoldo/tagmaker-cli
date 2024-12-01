package main

type Cell struct {
	widthPerUnit float64
	text         string
	title        bool
	style        string
}
type TagRow []Cell     // each element is a cell
type TagTable []TagRow // each element is a row with cells

type Tag struct { // cambiar model por Tag si vuelvo a GUI en vez de TUI
	width  float64 // in units of measurement
	height float64 // in unirs of measurement
	// lrMargin float64
	// tbMargin float64
	table TagTable // table shape metadata
}
