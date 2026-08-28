package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type InstructionsScreenModel struct {
	pages       []string
	currentPage int
	totalPages  int
	nextState   SessionState
}

func (m InstructionsScreenModel) Init() tea.Cmd {
	return nil
}

func InitialInstructionsScreenModel() InstructionsScreenModel {
	pages := []string{"InstructionsPage1", "InstructionsPage2", "InstructionsPage3"}
	return InstructionsScreenModel{
		pages:       pages,
		currentPage: 0,
		totalPages:  len(pages),
		nextState:   SessionStateInstructionsScreen,
	}
}

func (m *InstructionsScreenModel) GetNextState() SessionState {
	c := m.nextState
	m.nextState = SessionStateInstructionsScreen
	return c
}

func (m InstructionsScreenModel) Update(msg tea.Msg) (InstructionsScreenModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		switch msg.String() {
		case "q":
			m.nextState = SessionStateTitleScreen
			return m, nil
		case "left", "h":
			if m.currentPage > 0 {
				m.currentPage--
			}
		case "right", "l":
			if m.currentPage < m.totalPages-1 {
				m.currentPage++
			}
		}
	}
	return m, nil
}

func (m InstructionsScreenModel) View() tea.View {
	b := strings.Builder{}

	b.WriteString(banner)
	b.WriteString("\n")

	body := m.pages[m.currentPage]
	b.WriteString(boxStyle.Render(boxTextStyle.Render(body)))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Page %d of %d", m.currentPage+1, m.totalPages))
	b.WriteString("\n")
	b.WriteString(Footer)

	view := tea.NewView(b.String())
	return view
}
