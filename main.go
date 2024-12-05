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

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.currentView {
	case WelcomeView:
		return m.WelcomeKeys(msg)
	case TagView:
		if m.withTextInput {
			return m.TextInputKeys(msg)
		} else {
			return m.TagBuilderKeys(msg)
		}
	case BuildView:
	case FileLoaderView:
	case PrintView:
	}
	return m, nil
}

func (m *model) View() string {

	s := ""

	switch m.currentView {
	case WelcomeView:
		m.WelcomeView()
		s = fmt.Sprint(m.flexBox.Render())
	case TagView:
		if m.withTextInput {
			m.TagView(m.textInput.View())
		} else {
			m.TagView("")
		}
		s = fmt.Sprint(m.flexBox.Render())
	case BuildView:
	case FileLoaderView:
	case PrintView:
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

	csvData := CSVData{}

	csvData.headers = make([]string, 0)
	csvData.data = make([][]string, 0)

	csvData.headers, csvData.data, err = readCSVFile("tagdata.csv")

	if err != nil {
		log.Println(fmt.Errorf("%w", err))
	}

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
