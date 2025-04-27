package tui

import (
	"net/http"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipefbs/lazy-request/tui/keys"
)

var (
	modelStyle = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("238"))

	focusedModelStyle = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
)

type Model struct {
	width    int
	height   int
	focus    int
	tree     TreeSection
	request  *RequestSection
	response *ResponseSection
	list     []*http.Request
}

func New(list []*http.Request) Model {
	return Model{
		focus:    0,
		tree:     newTreeSection(list),
		request:  newRequestSection(),
		response: newResponseSection(),
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
			m.tree.Toggle()
		}
	}

	m.tree.SetFocus(m.focus == 0)
	m.request.SetFocus(m.focus == 1)
	m.response.SetFocus(m.focus == 2)

	var cmd tea.Cmd
	m.tree, cmd = m.tree.Update(msg)

	return m, cmd
}

func (m Model) RenderSection(curr int, strs ...string) string {
	if m.focus == curr {
		return focusedModelStyle.Render(strs...)
	} else {
		return modelStyle.Render(strs...)
	}
}

func (m Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		m.tree.View(),
		lipgloss.JoinVertical(lipgloss.Bottom,
			m.request.View(m),
			m.response.View(m),
		),
	)
}

type Screen struct {
	width   int
	height  int
	columns int
	rows    int
}
