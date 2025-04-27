package keys

import "github.com/charmbracelet/bubbles/key"

var (
	Up = key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	)
	Down = key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	)
	PrevPane = key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "Previus Pane"),
	)
	NextPane = key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "Next Pane"),
	)
	Select = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Select"),
	)
	Quit = key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "Quit"),
	)
	ToggleExplorer = key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "Toggle Explorer"),
	)
)
