package main

import (
	"fmt"
	"log"
	"os"

	"github.com/76creates/stickers/flexbox"
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
	// s += "\nArrows to move\tA: Insert row below\t\tZ: Delete current row\t\tX: Delete current cell\n\t\tS: Insert cell to the left\tD: Insert cell to the right\tB: Bind CSV data to Tag"
	// log.Print(s)
	return s
}

func (m *model) createRows(text string) {

	rows := []*flexbox.Row{}

	firstRow := m.flexBox.NewRow()
	firstRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
	rows = append(rows, firstRow)

	// Add tag rows
	for _, row := range m.tag.table {
		_fbRow := m.flexBox.NewRow()

		if _fbRow == nil {
			panic("could not find the table row")
		}
		// Add first padding cell before adding content cells
		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		if len(row) == 1 {
			row[0].widthPerUnit = 1.0
		}

		// Add content cells
		for j, cell := range row {
			style := m.cellStyleSelector(cell, styleNormal)

			_fbRow.AddCells(flexbox.NewCell(int(cell.widthPerUnit*100), 1).SetStyle(style))
			_cell := _fbRow.GetCell(j + 1).SetContent(cell.text) // +1 because of cell padding
			if _cell == nil {
				panic("could not find the table cell")
			}
		}
		// Add closing padding cell
		_fbRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))
		rows = append(rows, _fbRow)
	}
	// Add closing padding row
	lastRow := m.flexBox.NewRow()
	lastRow.AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG)).
		AddCells(flexbox.NewCell(100, 1).SetStyle(styleBG).SetContent(text)).
		AddCells(flexbox.NewCell(10, 1).SetStyle(styleBG))

	if m.textInputVisibility {
		lastRow.GetCell(1).SetStyle(styleTextInput)
	}
	rows = append(rows, lastRow)

	helpRow := m.flexBox.NewRow().AddCells(flexbox.NewCell(120, 5).SetStyle(styleHelp).
		SetContent("\nArrows to move\nA: Insert row below\t\tZ: Delete current row\t\tX: Delete current cell\nS: Insert cell to the left\tD: Insert cell to the right\tB: Bind CSV data to Tag fields"))

	rows = append(rows, helpRow)

	m.flexBox.SetRows(rows)

	// Highlight the current content row and cell as selected
	if (m.currCursorRow >= 0 && m.currCursorRow < len(m.tag.table)) &&
		(m.currCursorCell >= 0 && m.currCursorCell < len(m.tag.table[m.currCursorRow])) {

		cell := m.tag.table[m.currCursorRow][m.currCursorCell]
		style := m.cellStyleSelector(cell, styleSelected)
		rows[m.FBCursorRow()].GetCell(m.FBCursorCell()).SetStyle(style)
	}

	// SetRows instead of AddRows, since setrows overwrites, and when
	// calling CreateRows, we always want to overwrite to refresh the view.
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
