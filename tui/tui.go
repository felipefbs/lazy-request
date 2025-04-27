package tui

import (
	"net/http"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipefbs/lazy-request/tui/explorer"
	"github.com/felipefbs/lazy-request/tui/keys"
	"github.com/felipefbs/lazy-request/tui/request"
	"github.com/felipefbs/lazy-request/tui/response"
)

type Model struct {
	width    int
	height   int
	focus    int
	list     []*http.Request
	selected *http.Request
	explorer explorer.Explorer
	request  request.Request
	response response.Response
}

func New(list []*http.Request) Model {
	return Model{
		focus:    0,
		explorer: explorer.New(list),
		request:  request.New(list[1]),
		response: response.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.NextPane):
			m.focus = m.focus + 1
			if m.focus > 2 {
				m.focus = 0
			}
		case key.Matches(msg, keys.PrevPane):
			m.focus = m.focus - 1
			if m.focus < 0 {
				m.focus = 2
			}
		case key.Matches(msg, keys.ToggleExplorer):
			m.explorer.Toggle()
		}
	}

	m.explorer.SetFocus(m.focus == 0)
	m.request.SetFocus(m.focus == 1)
	m.response.SetFocus(m.focus == 2)

	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.explorer, cmd = m.explorer.Update(msg)
	cmds = append(cmds, cmd)

	m.request, cmd = m.request.Update(msg)
	cmds = append(cmds, cmd)

	m.response, cmd = m.response.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Sequence(cmds...)
}

func (m Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		m.explorer.View(),
		lipgloss.JoinVertical(lipgloss.Bottom,
			m.request.View(),
			m.response.View(),
		),
	)
}
