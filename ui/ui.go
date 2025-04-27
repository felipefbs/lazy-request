package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	width  int
	height int
	focus  int
	tree   *TreeModel
}

func NewModel() Model {
	return Model{
		focus: 0,
		tree:  newTreeModel(),
	}
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
		case "l":
			m.focus = m.focus + 1
			if m.focus > 2 {
				m.focus = 0
			}
		case "h":
			m.focus = m.focus - 1
			if m.focus < 0 {
				m.focus = 2
			}
		case "ctrl+s":
			m.tree.Toggle()
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height - 2
		m.width = msg.Width
		m.tree.SetDimention(m.height, m.height)
	}

	m.tree.SetFocus(m.focus == 0)

	return m, nil
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
		lipgloss.Top,
		m.tree.View(),
		lipgloss.JoinVertical(lipgloss.Top,
			m.RenderSection(1, "Request"),
			m.RenderSection(2, "Response"),
		),
	)
}

type Screen struct {
	width   int
	height  int
	columns int
	rows    int
}
