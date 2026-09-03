package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Drag0neUsz/Cipher-quest/internal/models"
)

func main() {
	args := os.Args
	if len(args) > 1 && args[1] == "-i" {
		p := tea.NewProgram(models.InitialInstructionsScreenModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	} else {
		p := tea.NewProgram(models.InitialRootModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}
}
