package models

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	content "github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type Cursor struct {
	chapter int
	puzzle  int
}

type Puzzle struct {
	Cipher      content.Cipher
	ID          string
	Title       string
	StateLink   SessionState
	IsCompleted bool
	IsLocked    bool
}

type Chapter struct {
	Puzzles []Puzzle
	Title   string
}

type ChapterSelectScreenModel struct {
	chapters  []Chapter
	cursor    Cursor
	nextState SessionState
}

func (m ChapterSelectScreenModel) Init() tea.Cmd {
	return nil
}

func InitialChapterSelectScreenModel() ChapterSelectScreenModel {
	chapters := []Chapter{
		{Title: "Chapter 1 Substitution Ciphers", Puzzles: []Puzzle{
			{ID: "caesar", Title: "Caesar Cipher", StateLink: SessionStatePuzzleScreen, IsCompleted: false, IsLocked: false, Cipher: content.CaesarCipher{Shift: 3}},
			{ID: "atbash", Title: "Atbash Cipher", StateLink: SessionStatePuzzleScreen, IsCompleted: false, IsLocked: false},
			{ID: "locked", Title: "Locked Cipher", StateLink: SessionStatePuzzleScreen, IsCompleted: false, IsLocked: true},
		}},
		{Title: "Chapter 2 Transposition Cipher", Puzzles: []Puzzle{
			{ID: "transposition", Title: "Transposition Cipher", StateLink: SessionStatePuzzleScreen, IsCompleted: false, IsLocked: false},
			{ID: "rail_fence", Title: "Rail Fence Cipher", StateLink: SessionStatePuzzleScreen, IsCompleted: false, IsLocked: true},
		}},
	}
	return ChapterSelectScreenModel{
		chapters:  chapters,
		cursor:    Cursor{chapter: 0, puzzle: 0},
		nextState: SessionStateChapterSelectScreen,
	}
}

func (m *ChapterSelectScreenModel) GetNextState() SessionState {
	c := m.nextState
	m.nextState = SessionStateChapterSelectScreen
	return c
}

func (m ChapterSelectScreenModel) GetSelectedPuzzle() Puzzle {
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

func (m ChapterSelectScreenModel) Update(msg tea.Msg) (ChapterSelectScreenModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		switch msg.String() {
		case "q":
			m.nextState = SessionStateTitleScreen
			return m, nil
		case "up", "k":
			m.cursor = m.handleCursorDecrement()

		case "down", "j":
			m.cursor = m.handleCursorIncrement(Cursor{chapter: m.cursor.chapter, puzzle: m.cursor.puzzle})

		case "enter", "space":
			m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].IsCompleted = !m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].IsCompleted
			m.nextState = m.chapters[m.cursor.chapter].Puzzles[m.cursor.puzzle].StateLink
		}
	}

	return m, nil
}

func (m ChapterSelectScreenModel) View() tea.View {
	b := strings.Builder{}
	b.WriteString(banner)
	b.WriteString("\n")

	for i, chapter := range m.chapters {
		b.WriteString(chapter.Title)
		b.WriteString("\n")
		for j, puzzle := range chapter.Puzzles {
			var row string
			if m.cursor.chapter == i && m.cursor.puzzle == j {
				row = selectedItemStyle.Render(fmt.Sprintf("%s %s", "->", puzzle.Title)) + "\n"
			} else {
				switch puzzle.IsLocked {
				case true:
					row = fmt.Sprintf("%s %s\n", "  ", lockedItemStyle.Render("\U0001F512", "Locked Puzzle"))
				case false:
					switch puzzle.IsCompleted {
					case true:
						row = fmt.Sprintf("%s %s\n", "  ", completedItemStyle.Render("[✓] ", puzzle.Title))
					case false:
						row = fmt.Sprintf("%s %s\n", "  ", incompleteItemStyle.Render("[ ] ", puzzle.Title))
					}
				}
			}
			b.WriteString(row)
		}
		b.WriteString("\n")
	}

	b.WriteString(Footer)

	view := tea.NewView(b.String())
	return view
}
