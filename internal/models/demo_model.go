package models

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type DemoModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func (m DemoModel) Init() tea.Cmd {
	return nil
}

func InitialDemoModel() DemoModel {
	return DemoModel{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m DemoModel) GetPreviousState() SessionState {
	return SessionStateTitleScreen
}

func (m DemoModel) GetNextState() SessionState {
	return SessionStateDemo
}

func (m DemoModel) Update(msg tea.Msg) (DemoModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m DemoModel) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// Send the UI for rendering
	return s
}
