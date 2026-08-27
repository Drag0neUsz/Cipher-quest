package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type TitleScreenModel struct {
	choices []string
	cursor  int
	choice  SessionState
}

func (m TitleScreenModel) Init() tea.Cmd {
	return nil
}

func InitialTitleScreenModel() TitleScreenModel {
	return TitleScreenModel{
		choices: []string{"Title Screen", "Demo"},
		cursor:  0,
		choice:  SessionStateTitleScreen,
	}
}

func (m TitleScreenModel) GetPreviousState() SessionState {
	return SessionStateTitleScreen
}

func (m *TitleScreenModel) GetNextState() SessionState {
	c := m.choice
	m.choice = SessionStateTitleScreen
	return c
}

func (m TitleScreenModel) Update(msg tea.Msg) (TitleScreenModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		// Cool, what was the actual key pressed?
		switch msg.String() {

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
			m.choice = SessionState(m.cursor)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m TitleScreenModel) View() string {
	b := strings.Builder{}
	// The header

	banner := `   ________ ____  __  _____________       ____  __  __________________
  / ____/ // __ \/ / / / ____/ __ \      / __ \/ / / / ____/ ___/_  __/
 / /   / // /_/ / /_/ / __/ / /_/ /_____/ / / / / / / __/  \__ \ / /   
/ /___/ // ____/ __  / /___/ _, _/_____/ /_/ / /_/ / /___ ___/ // /    
\____/_//_/   /_/ /_/_____/_/ |_|      \___\_\____/_____//____//_/

`
	b.WriteString(titleStyle.Render(banner))
	b.WriteString("\n")

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := "  " // no cursor
		if m.cursor == i {
			b.WriteString(selectedItemStyle.Render(fmt.Sprintf("%s %s", "->", choice)))
			b.WriteString("\n")
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

		// Render the row

	}

	// The footer
	b.WriteString("\nq to go back.    ctrl+c to quit.\n")

	// Send the UI for rendering
	return b.String()
}
