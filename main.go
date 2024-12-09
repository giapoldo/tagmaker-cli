package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// return textinput.Blink
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.currentView {
	case welcome1View:
		return m.welcomeKeys(msg)
	case welcome2View:
		return m.welcomeKeys(msg)
	case tagBuilderView:
		switch m.updateType {
		case normal:
			return m.tagBuilderKeys(msg)
		case textInput:
			return m.textInputKeys(msg)
		}
	case dataBinderView:
		return m.dataBindKeys(msg)
	case tagViewerView:
		switch m.updateType {
		case normal:
			return m.tagViewerKeys(msg)
		case textInput:
			return m.textInputKeys(msg)
		}
	case printToPDFView:
		switch m.updateType {
		case normal:
			return m.printToPDFKeys(msg)
		case textInput:
			return m.textInputKeys(msg)
		}
	}
	return m, nil
}

func (m *model) View() string {

	s := ""

	switch m.currentView {
	case welcome1View:
		m.welcome1View()
		s += fmt.Sprint(m.flexBox.Render())
	case welcome2View:
		m.welcome2View()
		s += fmt.Sprint(m.flexBox.Render())
	case tagBuilderView:
		switch m.updateType {
		case normal:
			m.tagBuilderView("")
		case textInput:
			m.tagBuilderView(m.textInput.View())
		}
		s += fmt.Sprint(m.flexBox.Render())

	case dataBinderView:
		if m.lastCSVHeaderIdx+1 == m.currentCSVHeaderIdx {
			m.dataBinderView(m.csvData.headers[m.currentCSVHeaderIdx])
		} else if m.lastCSVHeaderIdx == m.currentCSVHeaderIdx {
			m.dataBinderView(fmt.Sprintf("%s data", m.csvData.headers[m.currentCSVHeaderIdx]))
		}
		s += fmt.Sprint(m.flexBox.Render())

	case tagViewerView:
		switch m.updateType {
		case normal:
			m.tagViewerView("")
		case textInput:
			m.tagViewerView(m.textInput.View())
		}
		s += fmt.Sprint(m.flexBox.Render())
	case printToPDFView:

		m.printToPDFView(m.textInput.View())
		s += fmt.Sprint(m.flexBox.Render())
	}
	return s
}

func main() {

	// Setup file logging, since bubbletea hijacks the terminal and you can't print to screen.
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
