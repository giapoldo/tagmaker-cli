package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		windowHeight := msg.Height
		windowWidth := msg.Width
		m.flexBox.SetWidth(windowWidth)
		m.flexBox.SetHeight(windowHeight)

	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// // The "up" and "k" keys move the cursor up
		case "up":
			m.CursorUp()

		// // The "down" and "j" keys move the cursor down
		case "down":
			m.CursorDown()
			// // The "up" and "k" keys move the cursor up
		case "left":
			m.CursorLeft()

		// // The "down" and "j" keys move the cursor down
		case "right":
			m.CursorRight()

			// // The "enter" key and the spacebar (a literal space) toggle
			// // the selected state for the item that the cursor is pointing at.
			// case "enter", " ":
			// 	_, ok := m.selected[m.cursor]
			// 	if ok {
			// 		delete(m.selected, m.cursor)
			// 	} else {
			// 		m.selected[m.cursor] = struct{}{}
			// 	}
		}
	}

	return m, nil
}
func (m *model) View() string {
	m.flexBox.ForceRecalculate()

	// for i, row := range m.tag.table {
	// 	_fbRow := m.flexBox.GetRow(i + 1) // +1 because of row padding
	// 	if _fbRow == nil {
	// 		panic("could not find the table row")
	// 	}
	// 	for j, cell := range row {
	// 		_cell := _fbRow.GetCell(j + 1).SetContent(cell.text) // +1 because of cell padding
	// 		if _cell == nil {
	// 			panic("could not find the table cell")
	// 		}
	// 	}
	// }

	return m.flexBox.Render()
}

func main() {

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
