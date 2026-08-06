package games

import m "github.com/werener/2048.git/internal/core/models"

// EndlessGame is a game, which cannot be beaten
//
// It mostly follows same general rules of 2048, but now the [models.EndlessDistribution] is used
// Implements [models.Game] interface
type EndlessGame struct {
	grid  m.Grid
	score m.Score
}

// EndlessFromGame initializes an [EndlessGame], based on the previus state of the [models.Game]
func EndlessFromGame(ng m.Game) *EndlessGame {
	return &EndlessGame{
		grid:  *ng.Grid(),
		score: ng.Score(),
	}
}

// Score returns the current score of the game
func (game *EndlessGame) Score() m.Score {
	return game.score
}

// Grid returns the current grid state
func (game *EndlessGame) Grid() *m.Grid {
	return &game.grid
}

// State returns the current state of the game
func (game *EndlessGame) State() m.GameState {
	for _, direction := range []m.Direction{m.Up, m.Left, m.Down, m.Right} {
		if game.grid.CanShift(direction) {
			return m.RUNNING
		}
	}

	return m.DEFEAT
}

/*
MakeMove performs a full move cycle in the following order:

 1. Checks if a move in this direction will change the state of the board.
    If not - returns. Otherwise, executes steps 2-6.
 2. Shifts all tiles in the provided direction;
 3. Combines adjacent tiles with equal values;
 4. Shifts all tiles in the provided direction;
 5. Spawns a new random tile, using [models.Grid.SpawnTile] with [models.EndlessDistribution].
 6. Increments [EndlessGame.score] field by double the amount, gained from this move.
*/
func (game *EndlessGame) MakeMove(direction m.Direction) {
	if !game.grid.CanShift(direction) {
		return
	}
	game.score += 2 * game.grid.Shift(direction)

	game.grid.SpawnTile(m.EndlessDistribution)
}
