package models

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	content "github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type PuzzleScreenModel struct {
	puzzle                *content.Puzzle
	nextState             content.SessionState
	inputMode             bool
	currentInput          string
	lastInput             string
	textInput             textinput.Model
	submitAttempt         bool
	currentContentIndex   int
	completedContentIndex *int
	showStageCompletion   bool
}

func (m PuzzleScreenModel) GetUnlockQueue() []string {
	if m.puzzle.IsCompleted {
		return m.puzzle.UnlocksIDs
	}
	return []string{}
}

func (m PuzzleScreenModel) Init() tea.Cmd {
	return textinput.Blink
}

func InitialPuzzleScreenModel(puzzle *content.Puzzle) PuzzleScreenModel {
	ti := textinput.New()
	ti.Placeholder = "ENTER to start typing..."
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return PuzzleScreenModel{
		puzzle:                puzzle,
		nextState:             content.SessionStatePuzzleScreen,
		inputMode:             false,
		currentInput:          "",
		textInput:             ti,
		submitAttempt:         false,
		currentContentIndex:   0,
		completedContentIndex: &puzzle.CompletedContentIndex,
		showStageCompletion:   false,
	}
}

func checkAnswer(c content.Cipher, answer string, solution string) bool {
	return c.Encrypt(answer) == c.Encrypt(solution)
}

func (m PuzzleScreenModel) GetPreviousState() content.SessionState {
	return content.SessionStateTitleScreen
}

func (m *PuzzleScreenModel) GetNextState() content.SessionState {
	c := m.nextState
	m.nextState = content.SessionStatePuzzleScreen
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
				m.nextState = content.SessionStateChapterSelectScreen
				return m, nil

			case "enter", "space":
				if m.showStageCompletion {
					m.showStageCompletion = false
					if m.currentContentIndex < len(m.puzzle.Content)-1 {
						m.currentContentIndex++
					} else {
						m.puzzle.IsCompleted = true
						m.nextState = content.SessionStateChapterSelectScreen
					}
					return m, nil
				}
				if m.currentContentIndex > *m.completedContentIndex {
					m.inputMode = true
					m.currentInput = ""
				} else {
					m.showStageCompletion = true
				}
				m.textInput.Reset()
				m.submitAttempt = false
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
				result := checkAnswer(m.puzzle.Cipher, m.currentInput, m.puzzle.Content[m.currentContentIndex].Solution)
				if result {
					*m.completedContentIndex++
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

func (m PuzzleScreenModel) getProgressString() string {
	var completedString string
	if m.puzzle.IsCompleted {
		completedString = "[COMPLETED]"
	}
	return fmt.Sprintf(" (%d/%d) %s", m.currentContentIndex+1, len(m.puzzle.Content), completedString)
}

func (m PuzzleScreenModel) View() tea.View {
	b := strings.Builder{}
	b.WriteString(content.Banner)
	b.WriteString("\n")
	b.WriteString(m.puzzle.Title)
	b.WriteString(m.getProgressString())
	b.WriteString("\n")
	var body string
	if m.puzzle.Cipher != nil {
		if !m.showStageCompletion {
			body = m.puzzle.Content[m.currentContentIndex].Story
		} else {
			body = m.puzzle.Content[m.currentContentIndex].CompletionMessage
		}
	}
	b.WriteString(content.BoxStyle.Render(body))
	b.WriteString("\n")

	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += lipgloss.Height(b.String()) - 1
	}

	var submitText string
	if m.currentContentIndex <= *m.completedContentIndex {
		m.textInput.SetValue(m.puzzle.Content[m.currentContentIndex].Solution)
		submitText = "Great job! Press ENTER to continue..."
	} else if m.submitAttempt {
		submitText = "Not Quite! Try again!"
	}
	if !m.showStageCompletion {
		b.WriteString(m.textInput.View())
		b.WriteString("\n")
	}
	b.WriteString(submitText)

	b.WriteString("\n")
	b.WriteString(content.Footer)

	view := tea.NewView(b.String())
	if m.inputMode {
		view.Cursor = c
	}
	return view
}
