package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipefbs/lazy-request/store"
)

type Model struct {
	store   store.Store
	list    []string
	selectd string
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return "test"
}
