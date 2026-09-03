package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type TitleScreenModel struct {
	choices []string
	cursor  int
	choice  content.SessionState
}

var titleScreenStateChoices = []content.SessionState{
	content.SessionStateAboutScreen,
	content.SessionStateDemo,
	content.SessionStateChapterSelectScreen,
}

func (m TitleScreenModel) Init() tea.Cmd {
	return nil
}

func InitialTitleScreenModel() TitleScreenModel {
	return TitleScreenModel{
		choices: []string{"About", "Demo", "Chapter Select"},
		cursor:  0,
		choice:  content.SessionStateTitleScreen,
	}
}

func (m TitleScreenModel) GetPreviousState() content.SessionState {
	return content.SessionStateTitleScreen
}

func (m *TitleScreenModel) GetNextState() content.SessionState {
	c := m.choice
	m.choice = content.SessionStateTitleScreen
	return c
}

func (m TitleScreenModel) Update(msg tea.Msg) (TitleScreenModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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

func (m TitleScreenModel) View() tea.View {
	b := strings.Builder{}

	b.WriteString(content.Banner)
	b.WriteString("\n")

	for i, choice := range m.choices {

		cursor := "  "
		if m.cursor == i {
			b.WriteString(content.SelectedItemStyle.Render(fmt.Sprintf("%s %s", "->", choice)))
			b.WriteString("\n")
		} else {
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}

	}

	b.WriteString(content.Footer)

	view := tea.NewView(b.String())
	return view
}
