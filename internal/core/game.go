package core

type Score uint64
type GameState int

const (
	RUNNING GameState = iota
	DEFEAT
	WIN
)

// Game interface represents a set of methods, needed for displaying every visible aspect of a game of 2048
type Game interface {
	MakeMove(direction Direction)
	Score() (score Score)
	Grid() Grid
	State() GameState
}

// NormalGame is a game, which follows general rules of 2048
//
// Implements [Game] interface
type NormalGame struct {
	grid  Grid
	score Score
}

// NewNormalGame initializes a [NormalGame]
func NewNormalGame() *NormalGame {
	return &NormalGame{grid: NewGrid(4)}
}

// Score returns the current score of the game
func (game *NormalGame) Score() Score {
	return game.score
}

// Grid returns the current grid state
func (game *NormalGame) Grid() Grid {
	return game.grid
}

// state returns the current state of the game
func (game *NormalGame) State() GameState {
	for i := range game.grid.Size {
		for j := range game.grid.Size {
			if game.grid.Tiles[i][j].Value == 2048 {
				return WIN
			}
		}
	}

	for _, direction := range []Direction{Up, Left, Down, Right} {
		if game.grid.canShift(direction) {
			return RUNNING
		}
	}

	return DEFEAT
}

/*
MakeMove performs a full move cycle in the following order:

 1. Checks if a move in this direction will change the state of the board.
    If not - returns. Otherwise, executes steps (2-5).
 2. Shifts all tiles in the provided direction;
 3. Combines adjacent tiles with equal values;
 4. Shifts all tiles in the provided direction;
 5. Spawns a new random tile, using [Grid.SpawnTile] with [DefaultDistribution].
 6. Increments [NormalGame.score] field by the amount, gained from this move.
*/
func (game *NormalGame) MakeMove(direction Direction) {
	if !game.grid.canShift(direction) {
		return
	}
	game.score += game.grid.Shift(direction)

	game.grid.SpawnTile(DefaultDistribution)
}
