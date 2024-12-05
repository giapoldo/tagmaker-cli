package main

import tea "github.com/charmbracelet/bubbletea"

func (m *model) TextInputKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.inputTextValue = m.textInput.Value()
			m.UnSetUserInput()
			switch m.inputCaller {
			case leftInsert:
				m.InsertTagCellLeft()
			case rightInsert:
				m.InsertTagCellRight()
				// case "bind":
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *model) TagBuilderKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.SetCellInput(leftInsert)
		case "d":
			m.SetCellInput(rightInsert)

		}

	}

	return m, nil
}

func (m *model) WelcomeKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		default:
			m.currentView = TagView
		}
	}
	return m, nil
}
