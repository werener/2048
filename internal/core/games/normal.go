package games

import m "github.com/werener/2048.git/internal/core/models"

// NormalGame is a game, which follows general rules of 2048
//
// Implements [m.Game] interface
type NormalGame struct {
	grid  m.Grid
	score m.Score

	prevGrid  m.Grid
	prevScore m.Score
}

// NewNormalGame initializes a [NormalGame]
func NewNormalGame() *NormalGame {
	grid := m.NewGrid(4)
	grid.SpawnTile(m.DefaultDistribution)
	grid.SpawnTile(m.DefaultDistribution)

	return &NormalGame{grid: grid}
}

// Score returns the current score of the game
func (game *NormalGame) Score() m.Score {
	return game.score
}

// Grid returns the current grid state
func (game *NormalGame) Grid() *m.Grid {
	return &game.grid
}

// State returns the current state of the game
func (game *NormalGame) State() m.GameState {
	for i := range game.grid.Size {
		for j := range game.grid.Size {
			if game.grid.Tiles[i][j].Value == 2048 {
				return m.WIN
			}
		}
	}

	for _, direction := range []m.Direction{m.Up, m.Left, m.Down, m.Right} {
		if game.grid.CanShift(direction) {
			return m.RUNNING
		}
	}

	return m.DEFEAT
}

// Undo returns the game to its state as of last turn
func (game *NormalGame) Undo() {
	game.grid.Tiles = game.prevGrid.Tiles
	game.score = game.prevScore
}

/*
MakeMove performs a full move cycle in the following order:

 1. Checks if a move in this direction will change the state of the board.
    If not - returns. Otherwise, executes steps 2-7.
 2. Saves previous grid and scores for further use in [NormalGame.Undo]
 3. Shifts all tiles in the provided direction;
 4. Combines adjacent tiles with equal values;
 5. Shifts all tiles in the provided direction;
 6. Spawns a new random tile, using [Grid.SpawnTile] with [DefaultDistribution].
 7. Increments [NormalGame.score] field by the amount, gained from this move.
*/
func (game *NormalGame) MakeMove(direction m.Direction) {
	if !game.grid.CanShift(direction) {
		return
	}
	game.prevGrid = game.grid.Clone()
	game.prevScore = game.score

	game.score += game.grid.Shift(direction)
	game.grid.SpawnTile(m.DefaultDistribution)
}
