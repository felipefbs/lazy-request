package main

import (
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipefbs/lazy-request/tui"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	list := []*http.Request{
		req,
		req,
		req,
		req,
	}

	p := tea.NewProgram(tui.New(list), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
