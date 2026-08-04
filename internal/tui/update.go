package tui

import (
	_ "charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.grid.ShiftUp()
		case key.Matches(msg, m.keys.Left):
			m.grid.ShiftLeft()
		case key.Matches(msg, m.keys.Down):
			m.grid.ShiftDown()
		case key.Matches(msg, m.keys.Right):
			m.grid.ShiftRight()

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}
