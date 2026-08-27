package main

import (
	"fmt"
)

type Page interface {
	Render() string
	Update() Page
}

type MainPage struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func (p MainPage) Render() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range p.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if p.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := p.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	return s
}

func (p MainPage) Update() Page {
	return p
}
