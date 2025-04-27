package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipefbs/lazy-request/tui/styles"
)

type RequestSection struct {
	width  int
	height int
	focus  bool
}

func (s *RequestSection) SetFocus(focus bool) {
	s.focus = focus
}

func newRequestSection() RequestSection {
	return RequestSection{}
}

func (m RequestSection) Init() tea.Cmd {
	return nil
}

func (m RequestSection) Update(msg tea.Msg) (RequestSection, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m RequestSection) View() string {
	w, h := styles.Pane.GetFrameSize()
	height := (m.height / 2) - h

	width := m.width
	width -= int(float32(m.width)*0.3) - w
	width -= w * 2

	if m.focus {
		return lipgloss.NewStyle().
			Inherit(styles.FocusedPane).
			Height(height).
			Width(width).
			Render("Request")
	}

	return lipgloss.NewStyle().
		Inherit(styles.Pane).
		Height(height).
		Width(width).
		Render("Request")
}
