package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return textinput.Blink
	// return nil
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

		// // The the spacebar (a literal space) add a row to the tag
		case "r":
			m.InsertTagRow()
		case "d":
			m.DeleteTagRow()
		case "f":
			m.DeleteTagCell()
		case "c":
			/*err :=*/ m.InsertTagCellLeft(0.5)
			// if err != nil {
			// 	fmt.Println("fatal:", err)
			// 	os.Exit(1)
			// }

		case "v":
			/*err :=*/ m.InsertTagCellRight(0.5)
			// if err != nil {
			// 	fmt.Println("fatal:", err)
			// 	os.Exit(1)
			// }
		}

	}

	return m, nil
}
func (m *model) View() string {
	// m.flexBox.ForceRecalculate()
	return m.flexBox.Render()
}

func main() {

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
