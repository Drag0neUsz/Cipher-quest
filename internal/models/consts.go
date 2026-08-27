package models

import lipgloss "charm.land/lipgloss/v2"

type SessionState uint

const (
	SessionStateTitleScreen SessionState = iota
	SessionStateDemo
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF66")).
			Bold(true)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#00FF66")).
				Bold(true).
				Padding(0, 1)
)
