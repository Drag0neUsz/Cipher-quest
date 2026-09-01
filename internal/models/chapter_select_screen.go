package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type Cursor struct {
	chapter int
	puzzle  int
}

type ChapterSelectScreenModel struct {
	chapters    []content.Chapter
	cursor      Cursor
	nextState   content.SessionState
	unlockQueue []string
}

func (m ChapterSelectScreenModel) Init() tea.Cmd {
	return nil
}

func InitialChapterSelectScreenModel() ChapterSelectScreenModel {

	return ChapterSelectScreenModel{
		chapters:  content.Chapters,
		cursor:    Cursor{chapter: 0, puzzle: 0},
		nextState: content.SessionStateChapterSelectScreen,
	}
}

func (m *ChapterSelectScreenModel) GetNextState() content.SessionState {
	c := m.nextState
	m.nextState = content.SessionStateChapterSelectScreen
	return c
}

func (m ChapterSelectScreenModel) GetSelectedPuzzle() *content.Puzzle {
	return m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle]
}

// we don't want to have the cursor point at a locked puzzle, no need to check if 0:0 is locked because it wouldn't make sense for the first puzzle to be locked
func (m ChapterSelectScreenModel) handleCursorDecrement() Cursor {
	if m.cursor.puzzle > 0 {
		m.cursor.puzzle--
	} else if m.cursor.chapter > 0 {
		m.cursor.chapter--
		m.cursor.puzzle = len(m.chapters[m.cursor.chapter].Puzzles) - 1
	}
	if m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].IsLocked {
		return m.handleCursorDecrement()
	}
	return m.cursor
}

// we don't want to have the cursor point at a locked puzzle
func (m ChapterSelectScreenModel) handleCursorIncrement(lastOkay Cursor) Cursor {
	if m.cursor.puzzle == len(m.chapters[len(m.chapters)-1].Puzzles)-1 && m.cursor.chapter == len(m.chapters)-1 {
		return lastOkay
	}
	if m.cursor.puzzle < len(m.chapters[m.cursor.chapter].Puzzles)-1 {
		m.cursor.puzzle++
	} else if m.cursor.chapter < len(m.chapters)-1 {
		m.cursor.chapter++
		m.cursor.puzzle = 0
	}
	if m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].IsLocked {
		return m.handleCursorIncrement(lastOkay)
	}
	return m.cursor
}

func (m *ChapterSelectScreenModel) UnlockStuff() tea.Cmd {
	if len(m.unlockQueue) == 0 {
		return nil
	}
	for _, puzzleID := range m.unlockQueue {
	SearchPuzzle:
		for _, chapter := range m.chapters {
			for _, puzzle := range chapter.Puzzles {
				if puzzle.ID == puzzleID {
					puzzle.IsLocked = false
					break SearchPuzzle
				}
			}
		}
	}
	m.unlockQueue = []string{}
	return nil
}

func (m *ChapterSelectScreenModel) SetUnlockQueue(unlockQueue []string) {
	m.unlockQueue = unlockQueue
}

func (m ChapterSelectScreenModel) Update(msg tea.Msg) (ChapterSelectScreenModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		switch msg.String() {
		case "q":
			m.nextState = content.SessionStateTitleScreen
			return m, nil
		case "up", "k":
			m.cursor = m.handleCursorDecrement()

		case "down", "j":
			m.cursor = m.handleCursorIncrement(Cursor{chapter: m.cursor.chapter, puzzle: m.cursor.puzzle})

		case "enter", "space":
			// m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].IsCompleted = !m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].IsCompleted
			m.nextState = content.SessionStatePuzzleScreen
		}
	}

	return m, nil
}

func (m ChapterSelectScreenModel) View() tea.View {
	b := strings.Builder{}
	b.WriteString(content.Banner)
	b.WriteString("\n")

	for i, chapter := range m.chapters {
		b.WriteString(chapter.Title)
		b.WriteString("\n")
		for j, puzzle := range chapter.Puzzles {
			var row string
			if m.cursor.chapter == i && m.cursor.puzzle == j {
				row = content.SelectedItemStyle.Render(fmt.Sprintf("%s %s", "->", puzzle.Title)) + "\n"
			} else {
				switch puzzle.IsLocked {
				case true:
					row = fmt.Sprintf("%s %s\n", "  ", content.LockedItemStyle.Render("\U0001F512", "Locked Puzzle"))
				case false:
					switch puzzle.IsCompleted {
					case true:
						row = fmt.Sprintf("%s %s\n", "  ", content.CompletedItemStyle.Render("[✓] ", puzzle.Title))
					case false:
						row = fmt.Sprintf("%s %s\n", "  ", content.IncompleteItemStyle.Render("[ ] ", puzzle.Title))
					}
				}
			}
			b.WriteString(row)
		}
		b.WriteString("\n")
	}

	b.WriteString(content.Footer)

	view := tea.NewView(b.String())
	return view
}
