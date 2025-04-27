package styles

import "github.com/charmbracelet/lipgloss"

var (
	Pane = lipgloss.NewStyle().
		Align(lipgloss.Top, lipgloss.Top).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))

	FocusedPane = lipgloss.NewStyle().
			Inherit(Pane).
			BorderForeground(lipgloss.Color("69"))
)
