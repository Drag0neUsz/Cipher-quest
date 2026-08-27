package models

import lipgloss "charm.land/lipgloss/v2"

type SessionState uint

const (
	SessionStateTitleScreen SessionState = iota
	SessionStateDemo
	SessionStateAboutScreen
	SessionStateInstructionsScreen
	SessionStateChapterSelectScreen
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

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00AA44")).
			Padding(1, 2).
			Width(80)

	boxTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC"))

	footerStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("#555555"))

	lockedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444444"))

	completedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF66"))

	incompleteItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#666666"))
)

var banner = titleStyle.Render(`   ________ ____  __  _____________            ____  __  __________________
  / ____/ // __ \/ / / / ____/ __ \           / __ \/ / / / ____/ ___/_  __/
 / /   / // /_/ / /_/ / __/ / /_/ /  _____   / / / / / / / __/  \__ \ / /   
/ /___/ // ____/ __  / /___/ _, _/  _____   / /_/ / /_/ / /___ ___/ // /    
\____/_//_/   /_/ /_/_____/_/ |_|           \___\_\____/_____//____//_/

`)

var Footer = footerStyle.Render("\nq to go back.  •  ctrl+c to quit.")
