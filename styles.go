package main

import (
	"github.com/charmbracelet/lipgloss"
)

var styleNormal = lipgloss.NewStyle().
	Background(lipgloss.Color("#FFFFFF")).
	Foreground(lipgloss.Color("#000000")).
	Border(lipgloss.NormalBorder(), true, true, false, true). // each row handles their top border
	BorderBackground(lipgloss.Color("#FFFFFF")).
	BorderForeground(lipgloss.Color("#000000"))

var styleBG = lipgloss.NewStyle().Background(lipgloss.Color("#888888"))

var styleSelected = lipgloss.NewStyle().
	Background(lipgloss.Color("#43A5BE")).
	Foreground(lipgloss.Color("#000000")).
	Border(lipgloss.NormalBorder(), true, true, true, true). // each row handles their top border
	BorderBackground(lipgloss.Color("#43A5BE")).
	BorderForeground(lipgloss.Color("#000000"))
