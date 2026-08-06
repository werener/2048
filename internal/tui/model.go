package tui

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core"
)

var debugGrid = core.GridFromPreset(4, [][]uint64{
	{0, 2, 4, 8},
	{16, 64, 32, 0},
	{128, 256, 512, 0},
	{1024, 0, 2048, 0}})

type Size struct {
	width  int
	height int
}

type model struct {
	help help.Model // help component from Bubbles
	keys keyMap     // active application keybinds
	game core.Game  // current active game. Note: can be nil
	Size            // termial window size
}

func initialModel() model {
	help := help.New()

	return model{
		game: core.NewNormalGame(),
		help: help,
		keys: keys,
	}
}

func Run() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("Runtime error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
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
	Help: key.NewBinding(
		key.WithKeys("?", "/"),
		key.WithHelp(" [?]", "Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp(" [q]", "Quit"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}
func (k keyMap) FullHelp() [][]key.Binding {
	/*
		q w ?
		a s d
	*/
	return [][]key.Binding{
		{k.Quit, k.Left},  // first column
		{k.Up, k.Down},    // second column
		{k.Help, k.Right}, // third column
	}
}
