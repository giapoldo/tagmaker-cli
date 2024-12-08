package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) textInputKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.inputValue = m.textInput.Value()
			m.unsetUserInput()
			switch m.inputCaller {
			case cellLeftInsert:
				m.insertTagCellLeft()
			case cellRightInsert:
				m.insertTagCellRight()
			case changeCellWidth:
				m.changeCellWidth()
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *model) tagBuilderKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "b":
			m.currentView = dataBinderView
		case "up":
			m.tagCursorUp()
		case "down":
			m.tagCursorDown()
		case "left":
			m.tagCursorLeft()
		case "right":
			m.tagCursorRight()
		case "a":
			m.insertTagRow()
		case "z":
			m.deleteTagRow()
		case "x":
			m.deleteTagCell()
		case "s":
			m.setCellInput(cellLeftInsert)
		case "d":
			m.setCellInput(cellRightInsert)
		}

	}

	return m, nil
}

func (m *model) welcomeKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		// "A" only works on welcome2View
		case "a":
			if m.currentView == welcome1View {
				m.currentView = welcome2View

			} else {
				m.currentView = tagBuilderView
				m.insertTagRow()
			}
		default:
			m.currentView = welcome2View
		}
	}
	return m, nil
}

func (m *model) dataBindKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "ctrl+c", "q":
			return m, tea.Quit
		case "v":
			m.currentView = tagViewerView
		case " ":
			m.dataBindToCell()
		case "up":
			m.tagCursorUp()
		case "down":
			m.tagCursorDown()
		case "left":
			m.tagCursorLeft()
		case "right":
			m.tagCursorRight()
		case "backspace":
			m.skipBindToCell()
		case "esc":
			m.currentView = tagBuilderView
			m.updateType = normal
			m.currentCSVHeaderIdx = 0

		}
	}
	return m, cmd
}

func (m *model) tagViewerKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "ctrl+c", "q":
			return m, tea.Quit
		case "p":
			// create PDF
		case "b":
			m.toggleBold()
		case "i":
			m.toggleItalic()
		case "c":
			m.toggleCentered()
		case "w":
			// change cell width UNSAFE, you will need to set all cells in the row and make sure they add to 1.0
			m.setCellInput(changeCellWidth)
		case "k":
			m.previousTag()
		case "l":
			m.nextTag()
		case "up":
			m.tagCursorUp()
		case "down":
			m.tagCursorDown()
		case "left":
			m.tagCursorLeft()
		case "right":
			m.tagCursorRight()
		case "esc":
			m.currentView = tagBuilderView
			m.updateType = normal
			m.currentCSVHeaderIdx = 0

		}
	}
	return m, cmd
}
