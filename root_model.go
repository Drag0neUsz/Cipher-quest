package main

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/content"
	models "github.com/Drag0neUsz/Cipher-quest/internal/models"
)

type RootModel struct {
	state               content.SessionState
	titleScreen         models.TitleScreenModel
	aboutScreen         models.AboutScreenModel
	instructionsScreen  models.InstructionsScreenModel
	demo                models.DemoModel
	chapterSelectScreen models.ChapterSelectScreenModel
	puzzleScreen        models.PuzzleScreenModel
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func initialRootModel() RootModel {
	return RootModel{
		state:               content.SessionStateTitleScreen,
		titleScreen:         models.InitialTitleScreenModel(),
		instructionsScreen:  models.InitialInstructionsScreenModel(),
		demo:                models.InitialDemoModel(),
		chapterSelectScreen: models.InitialChapterSelectScreenModel(),
		aboutScreen:         models.InitialAboutScreenModel(),
	}
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	switch m.state {
	case content.SessionStateTitleScreen:
		m.titleScreen, cmd = m.titleScreen.Update(msg)
		if next := m.titleScreen.GetNextState(); next != m.state {
			m.state = next
		}

	case content.SessionStateAboutScreen:
		m.aboutScreen, cmd = m.aboutScreen.Update(msg)
		if next := m.aboutScreen.GetNextState(); next != m.state {
			m.state = next
		}

	case content.SessionStateInstructionsScreen:
		m.instructionsScreen, cmd = m.instructionsScreen.Update(msg)
		if next := m.instructionsScreen.GetNextState(); next != m.state {
			m.state = next
		}

	case content.SessionStateDemo:
		m.demo, cmd = m.demo.Update(msg)
		if next := m.demo.GetNextState(); next != m.state {
			m.state = next
		}

	case content.SessionStateChapterSelectScreen:
		m.chapterSelectScreen, cmd = m.chapterSelectScreen.Update(msg)

		// Obsługa wyboru konkretnego poziomu
		if next := m.chapterSelectScreen.GetNextState(); next != m.state {
			if next == content.SessionStatePuzzleScreen {
				// Dynamicznie ładujemy wybrany puzzle do ekranu gry
				puzzle := m.chapterSelectScreen.GetSelectedPuzzle()
				m.puzzleScreen = models.InitialPuzzleScreenModel(puzzle)
				m.state = next
				return m, m.puzzleScreen.Init()
			}
			m.state = next
		}

	case content.SessionStatePuzzleScreen:
		m.puzzleScreen, cmd = m.puzzleScreen.Update(msg)
		if next := m.puzzleScreen.GetNextState(); next != m.state {
			m.state = next
		}
	}

	return m, cmd
}

func (m RootModel) View() tea.View {
	var view tea.View
	switch m.state {
	case content.SessionStateTitleScreen:
		view = m.titleScreen.View()
	case content.SessionStateAboutScreen:
		view = m.aboutScreen.View()
	case content.SessionStateInstructionsScreen:
		view = m.instructionsScreen.View()
	case content.SessionStateDemo:
		view = m.demo.View()
	case content.SessionStateChapterSelectScreen:
		view = m.chapterSelectScreen.View()
	case content.SessionStatePuzzleScreen:
		view = m.puzzleScreen.View()
	}
	view.AltScreen = true
	return view

}
