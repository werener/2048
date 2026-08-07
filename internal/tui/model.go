package tui

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core/games"
	core "github.com/werener/2048.git/internal/core/models"
)

var debugGrid = core.GridFromPreset(4, [][]uint64{
	{0, 2, 4, 8},
	{16, 64, 32, 0},
	{128, 256, 512, 0},
	{1024, 0, 1024, 2048}})

type Size struct {
	width  int
	height int
}

type Popups struct {
	showEndlessPopup bool
}

type model struct {
	help      help.Model // help component from Bubbles
	keys      keyMap     // active application keybinds
	game      core.Game  // current active game. Note: can be nil
	undosLeft int        // shows the amount of times player can Undo last move
	Size                 // termial window size
	Popups               // determines what popups are currently shown on screen
}

func initialModel() model {
	help := help.New()
	game, undos := games.NewNormalGame(), 1
	// *game.Grid() = debugGrid
	return model{
		game:      game,
		undosLeft: undos,
		help:      help,
		keys:      keys,
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
