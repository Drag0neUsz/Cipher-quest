package models

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type AboutScreenModel struct {
	nextState SessionState
}

func (m AboutScreenModel) Init() tea.Cmd {
	return nil
}

func InitialAboutScreenModel() AboutScreenModel {
	return AboutScreenModel{
		nextState: SessionStateAboutScreen,
	}
}

func (m *AboutScreenModel) GetNextState() SessionState {
	c := m.nextState
	m.nextState = SessionStateAboutScreen
	return c
}

func (m AboutScreenModel) Update(msg tea.Msg) (AboutScreenModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		switch msg.String() {
		case "q":
			m.nextState = SessionStateTitleScreen
			return m, nil
		}
	}
	return m, nil
}

func (m AboutScreenModel) View() tea.View {
	b := strings.Builder{}
	// The header

	b.WriteString(banner)
	b.WriteString("\n")

	body := strings.Join([]string{
		"Welcome to Cipher-Quest!",
		"",
		"Cipher-Quest is a Crypto-Puzzle Game designed to fit right in your terminal. (You know, for that `cracking the code` vibe)",
		"",
		"You will be given a series of puzzles to solve, each will test your cryptographic knowledge.",
		"",
		"The puzzles will be in the form of cryptography problems, you will need to solve them in order to progress.",
		"",
		"Good luck! And don't forget to have fun!",
		"",
		"P.S. I may or may not have hidden some easter eggs in the code as well as a secret zoo level :D",
	}, "\n")
	b.WriteString(boxStyle.Render(boxTextStyle.Render(body)))

	b.WriteString(Footer)

	view := tea.NewView(b.String())
	return view
}
