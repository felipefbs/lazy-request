package ui

import "github.com/charmbracelet/lipgloss"

type TreeModel struct {
	width  int
	height int
	show   bool
	focus  bool
}

func newTreeModel() *TreeModel {
	return &TreeModel{
		show:  true,
		focus: true,
	}
}

func (m *TreeModel) SetDimention(w, h int) {
	m.width = w
	m.height = h
}

func (m *TreeModel) Toggle() {
	m.show = !m.show
}

func (m TreeModel) IsOpen() bool {
	return m.show
}

func (m *TreeModel) SetFocus(focus bool) {
	m.focus = focus
}

func (m TreeModel) View() string {
	if !m.show {
		return ""
	}

	if m.focus {
		return lipgloss.NewStyle().
			Height(m.height).
			Width(int(float32(m.width)*0.3)).
			Align(lipgloss.Top, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
			Render("Tree")
	}

	return lipgloss.NewStyle().
		Height(m.height).
		Width(int(float32(m.width)*0.3)).
		Align(lipgloss.Top, lipgloss.Left).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238")).
		Render("Tree")
}
