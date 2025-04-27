package tui

import (
	"net/http"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	request *http.Request
}

func (i item) Title() string {
	return i.request.Method
}

func (i item) Description() string {
	return i.request.URL.String()
}

func (i item) FilterValue() string {
	return i.request.URL.String()
}

type TreeSection struct {
	show   bool
	focus  bool
	height int
	width  int
	list   list.Model
}

func newTreeSection(reqs []*http.Request) TreeSection {
	items := []list.Item{}
	for _, r := range reqs {
		items = append(items, item{request: r})
	}

	return TreeSection{
		list:  list.New(items, list.NewDefaultDelegate(), 0, 0),
		show:  true,
		focus: true,
	}
}

func (m *TreeSection) Toggle() {
	m.show = !m.show
}

func (m TreeSection) IsOpen() bool {
	return m.show
}

func (m *TreeSection) SetFocus(focus bool) {
	m.focus = focus
}

func (m TreeSection) Init() tea.Cmd {
	return nil
}

func (m TreeSection) Update(msg tea.Msg) (TreeSection, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m TreeSection) View() string {
	if !m.show {
		return ""
	}

	w, h := pane.GetFrameSize()
	height := m.height - h
	width := int(float32(m.width)*0.3) - w

	p := lipgloss.NewStyle().
		Height(height).
		Width(width).
		MaxWidth(200)

	if m.focus {
		p = p.Inherit(focusedPane)
	} else {
		p = p.Inherit(pane)
	}

	return p.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		"Rezz",
		"Rezz",
	))
}
