package main

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/models"
)

type sessionState uint

const (
	sessionStateMain sessionState = iota
	sessionStateResults
)

type RootModel struct {
	state sessionState
	demo  models.DemoModel
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func initialRootModel() RootModel {
	return RootModel{
		state: sessionStateMain,
		demo:  models.InitialDemoModel(),
	}
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case sessionStateMain:
		m.demo, cmd = m.demo.Update(msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m RootModel) View() tea.View {
	switch m.state {
	case sessionStateMain:
		return m.demo.View()
	}

	return tea.NewView("Unknown state")
}
