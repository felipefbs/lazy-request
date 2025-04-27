package main

import (
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipefbs/lazy-request/store"
	"github.com/felipefbs/lazy-request/tui"
)

func main() {
	requests, err := store.ReadDirectory(".requests")
	if err != nil {
		log.Fatal(err)
	}

	list := []*http.Request{}
	for _, r := range requests {
		list = append(list, r.Request)
	}

	p := tea.NewProgram(tui.New(list), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
