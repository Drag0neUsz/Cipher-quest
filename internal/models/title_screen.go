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

var titleScreenStateChoices = []SessionState{
	SessionStateAboutScreen,
	SessionStateInstructionsScreen,
	SessionStateDemo,
	SessionStateChapterSelectScreen,
}

func (m TitleScreenModel) Init() tea.Cmd {
	return nil
}

func InitialTitleScreenModel() TitleScreenModel {
	return TitleScreenModel{
		choices: []string{"About", "Instructions", "Demo", "Chapter Select"},
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
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			m.choice = titleScreenStateChoices[m.cursor]
		}
	}

	return m, nil
}

func (m TitleScreenModel) View() string {
	b := strings.Builder{}

	b.WriteString(banner)
	b.WriteString("\n")

	for i, choice := range m.choices {

		cursor := "  "
		if m.cursor == i {
			b.WriteString(selectedItemStyle.Render(fmt.Sprintf("%s %s", "->", choice)))
			b.WriteString("\n")
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

	}

	return b.String()
}
