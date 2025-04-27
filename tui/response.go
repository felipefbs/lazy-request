package tui

import (
	"github.com/charmbracelet/lipgloss"
)

type ResponseSection struct {
	focus bool
}

func (s *ResponseSection) SetFocus(focus bool) {
	s.focus = focus
}

func newResponseSection() *ResponseSection {
	return &ResponseSection{}
}

func (s ResponseSection) View(m Model) string {
	height := m.height / 2

	width := m.width
	if m.tree.IsOpen() {
		width -= int(float32(m.width) * 0.3)
	}
	width -= 4

	if s.focus {
		return lipgloss.NewStyle().
			Height(height).
			Width(width).
			Align(lipgloss.Top, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
			Render("Response")
	}

	return lipgloss.NewStyle().
		Height(height).
		Width(width).
		Align(lipgloss.Top, lipgloss.Left).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238")).
		Render("Response")
}
