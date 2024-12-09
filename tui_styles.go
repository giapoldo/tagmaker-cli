package main

import (
	"log"

	"github.com/charmbracelet/lipgloss"
)

var styleNormal = lipgloss.NewStyle().
	Background(lipgloss.Color("#FFFFFF")).
	Foreground(lipgloss.Color("#000000")).
	Border(lipgloss.NormalBorder(), true, true, true, true). // each row handles their top border
	BorderBackground(lipgloss.Color("#FFFFFF")).
	BorderForeground(lipgloss.Color("#000000")).
	Bold(false).Italic(false).
	Align(lipgloss.Left, lipgloss.Center)

var styleBindList = lipgloss.NewStyle().
	Background(lipgloss.Color("#888888")).
	Foreground(lipgloss.Color("#000000")).
	Bold(false).Italic(false).
	Align(lipgloss.Center, lipgloss.Center)

var styleBG = lipgloss.NewStyle().Background(lipgloss.Color("#888888"))
var styleTextInput = lipgloss.NewStyle().
	Background(lipgloss.Color("#888888")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Align(lipgloss.Center, lipgloss.Center)

var stylePrintTextInput = lipgloss.NewStyle().
	Background(lipgloss.Color("#43A5BE")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Align(lipgloss.Center, lipgloss.Center)

var styleSelected = lipgloss.NewStyle().
	Background(lipgloss.Color("#43A5BE")).
	Foreground(lipgloss.Color("#000000")).
	Border(lipgloss.NormalBorder(), true, true, true, true). // each row handles their top border
	BorderBackground(lipgloss.Color("#43A5BE")).
	BorderForeground(lipgloss.Color("#000000")).
	Bold(false).Italic(false).
	Align(lipgloss.Left, lipgloss.Center)

var stylePermaSelected = lipgloss.NewStyle().
	Background(lipgloss.Color("#9D00FF")).
	Foreground(lipgloss.Color("#000000")).
	Border(lipgloss.NormalBorder(), true, true, true, true). // each row handles their top border
	BorderBackground(lipgloss.Color("#9D00FF")).
	BorderForeground(lipgloss.Color("#000000")).
	Bold(false).Italic(false).
	Align(lipgloss.Left, lipgloss.Center)

var styleWelcome = lipgloss.NewStyle().
	Background(lipgloss.Color("#43A5BE")).
	Foreground(lipgloss.Color("#000000")).
	Border(lipgloss.NormalBorder(), true, true, true, true). // each row handles their top border
	BorderBackground(lipgloss.Color("#43A5BE")).
	BorderForeground(lipgloss.Color("#000000")).
	Bold(false).Italic(false).
	Align(lipgloss.Center, lipgloss.Center)

var styleHelp = lipgloss.NewStyle().
	Bold(false).Italic(false).
	Align(lipgloss.Center, lipgloss.Center)

func (m *model) getCellTextStyle(cell cell, baseStyle lipgloss.Style) (style lipgloss.Style) {

	style = baseStyle

	if !cell.centered && cell.textStyle == "B" {
		style = baseStyle.Bold(true)

	} else if !cell.centered && cell.textStyle == "I" {
		style = baseStyle.Italic(true)

	} else if !cell.centered && cell.textStyle == "BI" {
		style = baseStyle.Bold(true).Italic(true)

	} else if cell.centered && cell.textStyle == "" {
		style = baseStyle.AlignHorizontal(lipgloss.Center)

	} else if cell.centered && cell.textStyle == "B" {
		style = baseStyle.Bold(true).AlignHorizontal(lipgloss.Center)

	} else if cell.centered && cell.textStyle == "I" {
		style = baseStyle.Italic(true).AlignHorizontal(lipgloss.Center)

	} else if cell.centered && cell.textStyle == "BI" {
		style = baseStyle.Bold(true).Italic(true).AlignHorizontal(lipgloss.Center)

	}

	return // default baseStyle is always regular font, left aligned and vertical centered
}

func (m *model) toggleBold() {

	cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]

	if cell.textStyle == "B" {
		cell.textStyle = ""
	} else if cell.textStyle == "I" {
		cell.textStyle = "BI"
	} else if cell.textStyle == "BI" {
		cell.textStyle = "I"
	} else if cell.textStyle == "" {
		cell.textStyle = "B"
	}

	m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = cell
	log.Print(m.tag.tagTable[m.tagRowCursor][m.tagCellCursor])
}

func (m *model) toggleItalic() {

	cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]

	if cell.textStyle == "I" {
		cell.textStyle = ""
	} else if cell.textStyle == "B" {
		cell.textStyle = "BI"
	} else if cell.textStyle == "BI" {
		cell.textStyle = "B"
	} else if cell.textStyle == "" {
		cell.textStyle = "I"
	}
	m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = cell

}

func (m *model) toggleCentered() {

	cell := m.tag.tagTable[m.tagRowCursor][m.tagCellCursor]

	cell.centered = !cell.centered

	m.tag.tagTable[m.tagRowCursor][m.tagCellCursor] = cell

}
