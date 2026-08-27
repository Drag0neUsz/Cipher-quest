package main

import (
	tea "charm.land/bubbletea/v2"
	models "github.com/Drag0neUsz/Cipher-quest/internal/models"
)

type RootModel struct {
	state       models.SessionState
	titleScreen models.TitleScreenModel
	demo        models.DemoModel
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func initialRootModel() RootModel {
	return RootModel{
		state:       models.SessionStateTitleScreen,
		titleScreen: models.InitialTitleScreenModel(),
		demo:        models.InitialDemoModel(),
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
			case models.SessionStateDemo:
				m.demo, cmd = m.demo.Update(msg)
				m.state = m.demo.GetNextState()
			}
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m RootModel) View() tea.View {
	var view tea.View
	switch m.state {
	case models.SessionStateTitleScreen:
		view = tea.NewView(m.titleScreen.View())
	case models.SessionStateDemo:
		view = tea.NewView(m.demo.View())
	}
	view.AltScreen = true
	return view

}
