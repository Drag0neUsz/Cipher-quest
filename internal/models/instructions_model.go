package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type InstructionsScreenModel struct {
	pages       []string
	currentPage int
	totalPages  int
	nextState   content.SessionState
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
		nextState:   content.SessionStateInstructionsScreen,
	}
}

func (m *InstructionsScreenModel) GetNextState() content.SessionState {
	c := m.nextState
	m.nextState = content.SessionStateInstructionsScreen
	return c
}

func (m InstructionsScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			m.nextState = content.SessionStateTitleScreen
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

	b.WriteString(content.Banner)
	b.WriteString("\n")

	body := m.pages[m.currentPage]
	b.WriteString(content.BoxStyle.Render(content.BoxTextStyle.Render(body)))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Page %d of %d", m.currentPage+1, m.totalPages))
	b.WriteString("\n")
	b.WriteString(content.FooterStyle.Render("ctrl+c to quit."))

	view := tea.View{
		Content:   b.String(),
		AltScreen: true,
	}
	return view
}
