package request

import (
	"context"
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipefbs/lazy-request/tui/styles"
)

type Request struct {
	width  int
	height int
	focus  bool
	req    *http.Request
}

func (s *Request) SetFocus(focus bool) {
	s.focus = focus
}

func New(req *http.Request) Request {
	return Request{
		req: req,
	}
}

func (m Request) Init() tea.Cmd {
	return nil
}

func (m Request) Update(msg tea.Msg) (Request, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m Request) View() string {
	height := (m.height - styles.Frame.Y) / 2

	width := m.width
	width -= int(float32(m.width)*0.3) - styles.Frame.X
	width -= styles.Frame.X * 2

	pane := lipgloss.NewStyle().
		Height(height).
		Width(width)

	if m.focus {
		pane = pane.Inherit(styles.FocusedPane)
	} else {
		pane = pane.Inherit(styles.Pane)
	}

	if m.req == nil {
		return pane.Render("Request")
	}

	body := "No body"
	r := m.req.Clone(context.Background())
	if r.Body != nil {
		b, err := io.ReadAll(r.Body)
		if err == nil {
			body = string(b)
		}
		defer r.Body.Close()
	}

	return pane.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.NewStyle().
				Width(width-styles.Frame.X).
				Render(
					lipgloss.JoinHorizontal(
						lipgloss.Top,
						styles.Pane.
							BorderStyle(lipgloss.HiddenBorder()).
							Bold(true).
							Background(lipgloss.Color("10")).
							Render(m.req.Method),
						styles.Pane.
							BorderStyle(lipgloss.HiddenBorder()).
							Render(m.req.URL.String()),
					),
				),
			styles.Divider.Width(width-styles.Frame.X).String(),
			styles.Pane.BorderStyle(lipgloss.HiddenBorder()).Render(body),
		),
	)
}
