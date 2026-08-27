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
	}
}

func (m InstructionsScreenModel) GetPreviousState() SessionState {
	return SessionStateTitleScreen
}

func (m *InstructionsScreenModel) GetNextState() SessionState {
	return SessionStateInstructionsScreen
}

func (m InstructionsScreenModel) Update(msg tea.Msg) (InstructionsScreenModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		switch msg.String() {
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

func (m InstructionsScreenModel) View() string {
	b := strings.Builder{}

	b.WriteString(banner)
	b.WriteString("\n")

	body := m.pages[m.currentPage]
	b.WriteString(boxStyle.Render(boxTextStyle.Render(body)))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Page %d of %d", m.currentPage+1, m.totalPages))

	return b.String()
}
