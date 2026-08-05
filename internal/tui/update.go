package tui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.score += m.grid.MakeMove(core.Up)
		case key.Matches(msg, m.keys.Left):
			m.score += m.grid.MakeMove(core.Left)
		case key.Matches(msg, m.keys.Down):
			m.score += m.grid.MakeMove(core.Down)
		case key.Matches(msg, m.keys.Right):
			m.score += m.grid.MakeMove(core.Right)
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}
