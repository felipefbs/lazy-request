package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type RequestSection struct {
	focus bool
}

func (s *RequestSection) SetFocus(focus bool) {
	s.focus = focus
}

func newRequestSection() *RequestSection {
	return &RequestSection{}
}

func (s RequestSection) View(m Model) string {
	height := (m.height / 2) - 1

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
			Render("Request")
	}

	return lipgloss.NewStyle().
		Height(height).
		Width(width).
		Align(lipgloss.Top, lipgloss.Left).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238")).
		Render("Request")
}
