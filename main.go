package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipefbs/lazy-request/ui"
)

func main() {
	p := tea.NewProgram(ui.NewModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
