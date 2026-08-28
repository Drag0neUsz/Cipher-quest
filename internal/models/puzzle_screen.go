package models

import (
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type PuzzleScreenModel struct {
	puzzle       Puzzle
	nextState    SessionState
	inputMode    bool
	currentInput string
	lastInput    string
	textInput    textinput.Model
}

func (m PuzzleScreenModel) Init() tea.Cmd {
	return textinput.Blink
}

func InitialPuzzleScreenModel(puzzle Puzzle) PuzzleScreenModel {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return PuzzleScreenModel{
		puzzle:       puzzle,
		nextState:    SessionStatePuzzleScreen,
		inputMode:    false,
		currentInput: "",
		lastInput:    "",
		textInput:    ti,
	}
}

func (m PuzzleScreenModel) GetPreviousState() SessionState {
	return SessionStateTitleScreen
}

func (m *PuzzleScreenModel) GetNextState() SessionState {
	c := m.nextState
	m.nextState = SessionStatePuzzleScreen
	return c
}

func (m PuzzleScreenModel) Update(msg tea.Msg) (PuzzleScreenModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		if !m.inputMode {
			switch msg.String() {
			case "q":
				m.nextState = SessionStateChapterSelectScreen
				return m, nil

			case "enter", "space":
				m.inputMode = true
				m.currentInput = ""
				return m, nil
			}
		} else {
			switch msg.String() {
			case "ctrl+c":
				m.inputMode = false
				m.currentInput = ""
				return m, tea.Quit
			case "enter":
				m.inputMode = false
				m.lastInput = m.currentInput
				m.currentInput = ""
				return m, nil
			default:
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}
		}
	}
	return m, cmd
}

func (m PuzzleScreenModel) View() tea.View {
	b := strings.Builder{}
	b.WriteString(banner)
	b.WriteString("\n")
	b.WriteString(m.puzzle.Title)
	b.WriteString("\n")
	if m.puzzle.Cipher != nil {
		b.WriteString(m.puzzle.Cipher.Description())
	}
	b.WriteString("\n")
	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += lipgloss.Height(b.String()) - 1
	}
	b.WriteString(m.textInput.View())
	b.WriteString("\n")

	b.WriteString(Footer)

	view := tea.NewView(b.String())
	if m.inputMode {
		view.Cursor = c
	}
	return view
}
