package explorer

import (
	"net/http"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipefbs/lazy-request/tui/styles"
)

type Explorer struct {
	show   bool
	focus  bool
	height int
	width  int
	list   list.Model
}

func New(reqs []*http.Request) Explorer {
	items := []list.Item{}
	for _, r := range reqs {
		items = append(items, item{request: r})
	}

	l := list.New(items, itemDelegate{}, 100, 10)
	l.Title = "Razz"
	l.SetShowStatusBar(false)
	l.SetShowFilter(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle

	return Explorer{
		list:  l,
		show:  true,
		focus: true,
	}
}

func (m *Explorer) Toggle() {
	m.show = !m.show
}

func (m Explorer) IsOpen() bool {
	return m.show
}

func (m *Explorer) SetFocus(focus bool) {
	m.focus = focus
}

func (m Explorer) Init() tea.Cmd {
	return nil
}

func (m Explorer) Update(msg tea.Msg) (Explorer, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	if m.focus {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Explorer) View() string {
	if !m.show {
		return ""
	}

	height := m.height - styles.Frame.Y
	width := int(float32(m.width)*0.3) - styles.Frame.X

	p := lipgloss.NewStyle().
		Height(height).
		Width(width).
		MaxWidth(200)

	if m.focus {
		p = p.Inherit(styles.FocusedPane)
	} else {
		p = p.Inherit(styles.Pane)
	}

	return p.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		m.list.View(),
	))
}
