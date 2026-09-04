package models

import (
	"fmt"
	"os"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/content"
)

type InstructionsScreenModel struct {
	pages       []string
	currentPage int
	totalPages  int
	nextState   content.SessionState
	viewport    viewport.Model
}

func (m InstructionsScreenModel) Init() tea.Cmd {
	return nil
}

func getInstructionsPages() []string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error gettingwd: %s", err)
		return []string{""}
	}
	file, err := os.ReadFile(wd + "/internal/content/instructions.md")
	if err != nil {
		fmt.Printf("error opening: %s", err)
		return []string{""}
	}
	raw := string(file)
	split := strings.Split(raw, "<!--pagebreak-->\n")
	return split
}

func InitialInstructionsScreenModel() InstructionsScreenModel {
	pages := getInstructionsPages()
	pager := viewport.New(viewport.WithWidth(74), viewport.WithHeight(8))
	pager.YPosition = 6
	pager.SetContent(pages[0])
	pager.SoftWrap = true
	return InstructionsScreenModel{
		pages:       pages,
		currentPage: 0,
		totalPages:  len(pages),
		nextState:   content.SessionStateInstructionsScreen,
		viewport:    pager,
	}
}

func (m *InstructionsScreenModel) GetNextState() content.SessionState {
	c := m.nextState
	m.nextState = content.SessionStateInstructionsScreen
	return c
}

func (m InstructionsScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			m.viewport.SetContent(m.pages[0])
			m.currentPage = 0
			m.viewport.SetYOffset(0)
			return m, nil
		case "left", "h":
			if m.currentPage > 0 {
				m.currentPage--
				m.viewport.SetContent(m.pages[m.currentPage])
				m.viewport.SetYOffset(0)
			}
		case "right", "l":
			if m.currentPage < m.totalPages-1 {
				m.currentPage++
				m.viewport.SetContent(m.pages[m.currentPage])
				m.viewport.SetYOffset(0)
			}
		}
	}
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m InstructionsScreenModel) View() tea.View {
	b := strings.Builder{}

	b.WriteString(content.Banner)
	b.WriteString("\n")

	body := m.viewport.View()
	b.WriteString(content.BoxStyle.Render(content.BoxTextStyle.Render(body)))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Page %d of %d", m.currentPage+1, m.totalPages))
	b.WriteString("\n")
	b.WriteString(content.FooterStyle.Render("q to go to start  •  ctrl+c to quit."))

	view := tea.View{
		Content:   b.String(),
		AltScreen: true,
		MouseMode: tea.MouseModeCellMotion,
	}
	return view
}
