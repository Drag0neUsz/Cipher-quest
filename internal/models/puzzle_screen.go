package models

import (
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	content "github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type PuzzleScreenModel struct {
	puzzle        *Puzzle
	nextState     SessionState
	inputMode     bool
	currentInput  string
	lastInput     string
	textInput     textinput.Model
	submitAttempt bool
}

func (m PuzzleScreenModel) Init() tea.Cmd {
	return textinput.Blink
}

func InitialPuzzleScreenModel(puzzle *Puzzle) PuzzleScreenModel {
	ti := textinput.New()
	ti.Placeholder = "ENTER to start typing..."
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return PuzzleScreenModel{
		puzzle:        puzzle,
		nextState:     SessionStatePuzzleScreen,
		inputMode:     false,
		currentInput:  "",
		textInput:     ti,
		submitAttempt: false,
	}
}

func checkAnswer(c content.Cipher, answer string, solution string) bool {
	return c.Encrypt(answer) == c.Encrypt(solution)
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
				m.textInput.Reset()
				return m, nil
			}
			m.submitAttempt = false
		} else {
			switch msg.String() {
			case "ctrl+c":
				m.inputMode = false
				return m, tea.Quit
			case "enter":
				m.submitAttempt = true
				m.inputMode = false
				m.currentInput = m.textInput.Value()
				result := checkAnswer(m.puzzle.Cipher, m.currentInput, m.puzzle.Solution)
				if result {
					m.puzzle.IsCompleted = true
				}
				return m, nil
			default:
				m.textInput, cmd = m.textInput.Update(msg)
				m.submitAttempt = false
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
	body := strings.Builder{}
	if m.puzzle.Cipher != nil {
		body.WriteString(m.puzzle.Cipher.Description())
	}
	body.WriteString("\n")
	b.WriteString(boxStyle.Render(body.String()))
	b.WriteString("\n")

	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += lipgloss.Height(b.String()) - 1
	}

	b.WriteString(m.textInput.View())
	b.WriteString("\n")
	if m.puzzle.IsCompleted {
		b.WriteString("Congratulations! You have completed this puzzle!")
	} else if m.submitAttempt {
		b.WriteString("Not Quite! Try again!")
	}
	b.WriteString("\n")

	b.WriteString(Footer)

	view := tea.NewView(b.String())
	if m.inputMode {
		view.Cursor = c
	}
	return view
}
