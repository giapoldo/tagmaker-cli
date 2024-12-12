package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

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
			m.tagRowCursor = 0
			m.tagCellCursor = 0
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
			m.setUserInput(cellLeftInsert)
		case "d":
			m.setUserInput(cellRightInsert)
		}

	}

	return m, nil
}

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
			case setTagSize, setFontSize:
				switch m.printRowCursor {
				case 2:
					m.tag.width = m.getCellSizeValue()
				case 3:
					m.tag.height = m.getCellSizeValue()
				case 4:
					m.tag.fontSize = m.getCellSizeValue()
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
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
			m.tagRowCursor = 0
			m.tagCellCursor = 0
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
			m.tagRowCursor = 0
			m.tagCellCursor = 0

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
			m.currentView = printToPDFView
			m.printRowCursor = 1
			m.printCellCursor = 1
		case "b":
			m.toggleBold()
		case "i":
			m.toggleItalic()
		case "c":
			m.toggleCentered()
		case "w":
			// change cell width UNSAFE, you will need to set all cells in the row and make sure they add to 1.0
			m.setUserInput(changeCellWidth)
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
			m.currentView = dataBinderView
			m.updateType = normal
			m.currentCSVHeaderIdx = 0
			m.tagRowCursor = 0
			m.tagCellCursor = 0

		}
	}
	return m, cmd
}

func (m *model) printToPDFKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			if m.printCellCursor >= 2 {
				if m.printRowCursor == 1 {
					m.getPaperSize()
				}
			}
		case " ":
			if m.printCellCursor >= 2 {
				switch m.printRowCursor {
				case 2, 3, 4:
					m.setUserInput(setTagSize)
				}
			}
		case "p":
			m.pdfGenerator()
		case "up":
			m.printCursorUp()
		case "down":
			m.printCursorDown()
		case "left":
			m.printCursorLeft()
		case "right":
			m.printCursorRight()
		case "esc":
			m.currentView = tagViewerView
			m.updateType = normal
			m.currentCSVHeaderIdx = 0
			m.tagRowCursor = 0
			m.tagCellCursor = 0

		}
	}
	return m, cmd
}
