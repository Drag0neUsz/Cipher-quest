package content

import lipgloss "charm.land/lipgloss/v2"

type SessionState uint

const (
	SessionStateTitleScreen SessionState = iota
	SessionStateDemo
	SessionStateAboutScreen
	SessionStateInstructionsScreen
	SessionStateChapterSelectScreen
	SessionStatePuzzleScreen
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF66")).
			Bold(true)

	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#00FF66")).
				Bold(true).
				Padding(0, 1)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00AA44")).
			Padding(1, 2).
			Width(80)

	BoxTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC"))

	FooterStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("#555555"))

	LockedItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444444"))

	CompletedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF66"))

	IncompleteItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#666666"))

	TagStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#00AA44")).
			Foreground(lipgloss.Color("#000000")).
			Bold(true).
			Padding(0, 1)

	QuoteStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color("#00FF66")).
			Foreground(lipgloss.Color("#CCCCCC")).
			PaddingLeft(2).
			Italic(true)

	QuoteTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC")).
			Italic(true)

	QuoteAccent = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF66")).
			Bold(true)

	HintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	PayloadTagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF66")).
			Bold(true)

	PayloadBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00AA55")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2)
)

var Banner = TitleStyle.Render(`   ________ ____  __  _____________            ____  __  __________________
  / ____/ // __ \/ / / / ____/ __ \           / __ \/ / / / ____/ ___/_  __/
 / /   / // /_/ / /_/ / __/ / /_/ /  _____   / / / / / / / __/  \__ \ / /   
/ /___/ // ____/ __  / /___/ _, _/  _____   / /_/ / /_/ / /___ ___/ // /    
\____/_//_/   /_/ /_/_____/_/ |_|           \___\_\____/_____//____//_/

`)

var Footer = FooterStyle.Render("\nq to go back.  •  ctrl+c to quit.")
