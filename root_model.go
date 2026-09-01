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
	var next content.SessionState
	switch m.state {
	case content.SessionStateTitleScreen:
		m.titleScreen, cmd = m.titleScreen.Update(msg)
		next = m.titleScreen.GetNextState()

	case content.SessionStateAboutScreen:
		m.aboutScreen, cmd = m.aboutScreen.Update(msg)
		next = m.aboutScreen.GetNextState()

	case content.SessionStateInstructionsScreen:
		m.instructionsScreen, cmd = m.instructionsScreen.Update(msg)
		next = m.instructionsScreen.GetNextState()

	case content.SessionStateDemo:
		m.demo, cmd = m.demo.Update(msg)
		next = m.demo.GetNextState()

	case content.SessionStateChapterSelectScreen:
		m.chapterSelectScreen, cmd = m.chapterSelectScreen.Update(msg)

		next = m.chapterSelectScreen.GetNextState()
		if next == content.SessionStatePuzzleScreen {
			puzzle := m.chapterSelectScreen.GetSelectedPuzzle()
			m.puzzleScreen = models.InitialPuzzleScreenModel(puzzle)
			m.state = content.SessionStatePuzzleScreen
			return m, m.puzzleScreen.Init()
		}

	case content.SessionStatePuzzleScreen:
		m.puzzleScreen, cmd = m.puzzleScreen.Update(msg)
		next = m.puzzleScreen.GetNextState()
		if next == content.SessionStateChapterSelectScreen {
			m.state = next
			m.chapterSelectScreen.SetUnlockQueue(m.puzzleScreen.GetUnlockQueue())
			return m, m.chapterSelectScreen.UnlockStuff()
		}

	}

	if next != m.state {
		m.state = next
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
