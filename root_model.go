package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	models "github.com/Drag0neUsz/Cipher-quest/internal/models"
)

type RootModel struct {
	state               models.SessionState
	titleScreen         models.TitleScreenModel
	aboutScreen         models.AboutScreenModel
	instructionsScreen  models.InstructionsScreenModel
	demo                models.DemoModel
	chapterSelectScreen models.ChapterSelectScreenModel
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func initialRootModel() RootModel {
	return RootModel{
		state:               models.SessionStateTitleScreen,
		titleScreen:         models.InitialTitleScreenModel(),
		instructionsScreen:  models.InitialInstructionsScreenModel(),
		demo:                models.InitialDemoModel(),
		chapterSelectScreen: models.InitialChapterSelectScreenModel(),
	}
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case tea.KeyMsg:
		msg := msg.(tea.KeyMsg)
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			m.state = m.titleScreen.GetPreviousState()
			return m, nil

		default:
			switch m.state {
			case models.SessionStateTitleScreen:
				m.titleScreen, cmd = m.titleScreen.Update(msg)
				m.state = m.titleScreen.GetNextState()
			case models.SessionStateAboutScreen:
				m.aboutScreen, cmd = m.aboutScreen.Update(msg)
				m.state = m.aboutScreen.GetNextState()
			case models.SessionStateInstructionsScreen:
				m.instructionsScreen, cmd = m.instructionsScreen.Update(msg)
				m.state = m.instructionsScreen.GetNextState()
			case models.SessionStateDemo:
				m.demo, cmd = m.demo.Update(msg)
				m.state = m.demo.GetNextState()
			case models.SessionStateChapterSelectScreen:
				m.chapterSelectScreen, cmd = m.chapterSelectScreen.Update(msg)
				m.state = m.chapterSelectScreen.GetNextState()
			}
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m RootModel) View() tea.View {
	view := strings.Builder{}
	switch m.state {
	case models.SessionStateTitleScreen:
		view.WriteString(m.titleScreen.View())
	case models.SessionStateAboutScreen:
		view.WriteString(m.aboutScreen.View())
	case models.SessionStateInstructionsScreen:
		view.WriteString(m.instructionsScreen.View())
	case models.SessionStateDemo:
		view.WriteString(m.demo.View())
	case models.SessionStateChapterSelectScreen:
		view.WriteString(m.chapterSelectScreen.View())
	}
	view.WriteString(models.Footer)
	teaView := tea.NewView(view.String())
	teaView.AltScreen = true
	return teaView

}
