package response

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipefbs/lazy-request/tui/styles"
)

type Response struct {
	width  int
	height int
	focus  bool
}

func (s *Response) SetFocus(focus bool) {
	s.focus = focus
}

func New() Response {
	return Response{}
}

func (m Response) Init() tea.Cmd {
	return nil
}

func (m Response) Update(msg tea.Msg) (Response, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m Response) View() string {
	w, h := styles.Pane.GetFrameSize()
	height := (m.height - h*2) / 2

	width := m.width
	width -= int(float32(m.width)*0.3) - w
	width -= w * 2

	if m.focus {
		return lipgloss.NewStyle().
			Inherit(styles.FocusedPane).
			Height(height).
			Width(width).
			Render("Response")
	}

	return lipgloss.NewStyle().
		Inherit(styles.Pane).
		Height(height).
		Width(width).
		Render("Response")
}
