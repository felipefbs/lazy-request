package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Pane = lipgloss.NewStyle().
		Align(lipgloss.Top, lipgloss.Top).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))

	FocusedPane = lipgloss.NewStyle().
			Inherit(Pane).
			BorderForeground(lipgloss.Color("69"))

	Divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(lipgloss.Color("69"))
)

type frame struct {
	X, Y int
}

var Frame = frame{
	X: Pane.GetHorizontalFrameSize(),
	Y: Pane.GetVerticalFrameSize(),
}
