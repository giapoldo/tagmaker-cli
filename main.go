package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// return textinput.Blink
	return nil
}

func (m *model) TextInputOperations(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
		case "enter":
			m.textValue = m.textInput.Value()
			m.UnSetUserInput()
			switch m.inputCaller {
			case "left":
				m.InsertTagCellLeft()
			case "right":
				m.InsertTagCellRight()
			case "bind":
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *model) TagOperations(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "up":
			m.CursorUp()
		case "down":
			m.CursorDown()
		case "left":
			m.CursorLeft()
		case "right":
			m.CursorRight()
		case "a":
			m.InsertTagRow()
		case "z":
			m.DeleteTagRow()
		case "x":
			m.DeleteTagCell()
		case "s":
			m.SetCellInput("left")
		case "d":
			m.SetCellInput("right")
		}

	}

	return m, nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.textInputVisibility {
		return m.TextInputOperations(msg)
	} else {
		return m.TagOperations(msg)
	}
}

func (m *model) View() string {

	if m.textInputVisibility {
		m.createRows(m.textInput.View())
	} else {
		m.createRows("")
	}
	s := fmt.Sprint(m.flexBox.Render())
	s += "\nArrows to move\tA: Insert row below\t\tZ: Delete current row\t\tX: Delete current cell\n\t\tS: Insert cell to the left\tD: Insert cell to the right"
	log.Print(s)
	return s
}

func main() {

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
