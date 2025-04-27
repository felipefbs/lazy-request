package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type TreeSection struct {
	show  bool
	focus bool
}

func newTreeSection() *TreeSection {
	return &TreeSection{
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

func (m TreeSection) View(width, height int) string {
	if !m.show {
		return ""
	}

	if m.focus {
		return lipgloss.NewStyle().
			Height(height).
			Width(int(float32(width)*0.3)).
			Align(lipgloss.Top, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
			Render("Tree")
	}

	return lipgloss.NewStyle().
		Height(height).
		Width(int(float32(width)*0.3)).
		Align(lipgloss.Top, lipgloss.Left).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238")).
		Render("Tree")
}
