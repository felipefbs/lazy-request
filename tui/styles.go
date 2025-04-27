package tui

import "github.com/charmbracelet/lipgloss"

var (
	pane = lipgloss.NewStyle().
		Align(lipgloss.Top, lipgloss.Top).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))

	focusedPane = lipgloss.NewStyle().
			Inherit(pane).
			BorderForeground(lipgloss.Color("69"))
)
