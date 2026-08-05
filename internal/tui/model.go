package tui

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core"
)

type GameState int

const (
	RUNNING GameState = iota
	DEFEAT
	DRAW
	WIN
)

type Size struct {
	width  int
	height int
}
type Game struct {
	state GameState // in what state the application is currently in
	grid  core.Grid // grid, containing current state of the game
	score uint64    // current score
}

var debugGrid = core.GridFromPreset(4, [][]uint64{
	{0, 2, 4, 8},
	{16, 64, 32, 0},
	{128, 256, 512, 0},
	{1024, 0, 2048, 0}})

func newGame() Game {
	return Game{state: RUNNING,
		grid: debugGrid,
		// grid:  core.NewGrid(4),
		score: 0,
	}
}

type model struct {
	help help.Model
	keys keyMap // active application keybinds
	Game        // fields, related to the game runtime
	Size        // termial window size
}

func initialModel() model {
	help := help.New()

	return model{
		Game: newGame(),
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
		{k.Right, k.Help}, // third column
	}
}
