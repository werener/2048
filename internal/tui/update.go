package tui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core/games"
	core "github.com/werener/2048.git/internal/core/models"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	gameState := m.game.State()

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		// always allow to expand Help, Restart and Quit
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Restart):
			m.game = games.NewNormalGame()
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

		// read move keys only if the game is active
		switch gameState {
		case core.RUNNING:
			return m.updateRunning(msg)
		case core.WIN:
			return m.updateWin(msg)
		}
	default:
		if m.showEndlessPopup {
			m.showEndlessPopup = false
			m.game = games.EndlessFromGame(m.game)
		}
	}
	return m, nil
}

func (m model) updateRunning(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyPressMsg); ok {
		switch {
		case key.Matches(msg, m.keys.Up):
			m.game.MakeMove(core.Up)
		case key.Matches(msg, m.keys.Left):
			m.game.MakeMove(core.Left)
		case key.Matches(msg, m.keys.Down):
			m.game.MakeMove(core.Down)
		case key.Matches(msg, m.keys.Right):
			m.game.MakeMove(core.Right)
		case key.Matches(msg, m.keys.Undo):
			ng, isNormalGame := (m.game).(*games.NormalGame)
			if isNormalGame && m.undosLeft > 0 {
				ng.Undo()
				m.undosLeft--
			}
		}
	}
	return m, nil
}

func (m model) updateWin(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyPressMsg); ok {
		switch {
		case key.Matches(msg, m.keys.Continue):
			if m.showEndlessPopup {
				m.showEndlessPopup = false
				m.game = games.EndlessFromGame(m.game)
			} else {
				m.showEndlessPopup = true
			}
		}
	}
	return m, nil
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Undo  key.Binding

	Help     key.Binding
	Quit     key.Binding
	Restart  key.Binding
	Continue key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k", "w"),
		key.WithHelp("w|↑|k", "Up"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h", "a"),
		key.WithHelp("a|←|h", "Left"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j", "s"),
		key.WithHelp("s|↓|j", "Down"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l", "d"),
		key.WithHelp("d|→|l", "Right"),
	),
	Undo: key.NewBinding(
		key.WithKeys("ctrl+z", "u", "del"),
		key.WithHelp("u|Ctrl+Z|Del", "Undo"),
	),

	Help: key.NewBinding(
		key.WithKeys("?", "/"),
		key.WithHelp("[?]", "Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "Q", "esc", "ctrl+c"),
		key.WithHelp("[q]", "Quit"),
	),
	Restart: key.NewBinding(
		key.WithKeys("r", "R"),
		key.WithHelp("[r]", "Restart"),
	),

	Continue: key.NewBinding(
		key.WithKeys("c", "C"),
		key.WithHelp("[c]", "Continue"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Restart, k.Help}
}
func (k keyMap) FullHelp() [][]key.Binding {
	/*
		q w r
		a s d
	*/
	return [][]key.Binding{
		{k.Quit, k.Left},       // first column
		{k.Up, k.Down, k.Undo}, // second column
		{k.Restart, k.Right},   // third column
	}
}
