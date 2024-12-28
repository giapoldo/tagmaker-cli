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
			m.saveInputValue()
			m.inputCaller()
			m.unsetUserInput()

		case "up":
			if m.currentView == printToPDFView {
				m.printCursorUp()
			} else {
				m.tagCursorUp()
			}
			m.unsetUserInput()
		case "down":
			if m.currentView == printToPDFView {
				m.printCursorDown()
			} else {
				m.tagCursorDown()
			}
			m.unsetUserInput()
		case "left":
			if m.currentView == printToPDFView {
				m.printCursorLeft()
			} else {
				m.tagCursorLeft()
			}
			m.unsetUserInput()
		case "right":
			if m.currentView == printToPDFView {
				m.printCursorRight()
			} else {
				m.tagCursorRight()
			}
			m.unsetUserInput()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *model) tagKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "up":
			m.tagCursorUp()
		case "down":
			m.tagCursorDown()
		case "left":
			m.tagCursorLeft()
		case "right":
			m.tagCursorRight()
		}

		switch m.currentView {
		case tagBuilderView:
			switch msg.String() {
			case "a":
				m.insertTagRow()
			case "x":
				m.deleteTagCell()
			case "s":
				m.setUserInput(m.insertTagCellLeft)
				m.textInput.Placeholder = "Enter width in per unit (0.20~0.80)"
			case "d":
				m.setUserInput(m.insertTagCellRight)
				m.textInput.Placeholder = "Enter width in per unit (0.20~0.80)"
			case "n":
				m.currentView = dataBinderView
				m.resetViewState()
			case "esc":
				m.resetViewState()
			}
		case dataBinderView:
			switch msg.String() {
			case " ":
				m.dataBindToCell()
			case "backspace":
				m.skipBindToCell()
			case "n":
				m.currentView = tagViewerView
				m.resetViewState()
			case "esc":
				m.currentView = tagBuilderView
				m.resetViewState()
			}
		case tagViewerView:
			switch msg.String() {
			case "b":
				m.toggleBold()
			case "i":
				m.toggleItalic()
			case "c":
				m.toggleCentered()
			case "w":
				// change cell width UNSAFE, you will need to set all cells in the row and make sure they add to 1.0
				m.setUserInput(m.changeCellWidth)
				m.textInput.Placeholder = "Enter width in per unit (0.20~0.80)"
			case "k":
				m.previousTag()
			case "l":
				m.nextTag()
			case "n":
				m.currentView = printToPDFView
				m.resetViewState()
			case "esc":
				m.currentView = dataBinderView
				m.resetViewState()
			}
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
			switch printViewRows[m.printRowCursor] {
			case printViewRows[1]: // paper size
				m.pVContents.selectedValues[printKeysStatic[1]] = m.getPaperSize()

			case printViewRows[2]: // width
				m.setUserInput(m.saveToCurrentPVSelected)
				m.textInput.Placeholder = "Enter width"

			case printViewRows[3]: // height
				m.setUserInput(m.saveToCurrentPVSelected)
				m.textInput.Placeholder = "Enter height"

			case printViewRows[4]: // font size
				m.setUserInput(m.saveToCurrentPVSelected)
				m.textInput.Placeholder = "Enter font size"

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
			m.resetViewState()

		}
	}
	return m, cmd
}
