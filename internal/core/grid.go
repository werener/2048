package core

import (
	"math/rand"
	"slices"
)

type usize uint16

type Direction int8

const (
	Up Direction = iota
	Left
	Down
	Right
)

var DefaultDistribution = []uint64{4, 2, 2, 2, 2, 2, 2, 2, 2, 2}

// Grid represents the aggregation of playing blocks.
type Grid struct {
	Size  usize
	Tiles [][]Tile
}

// GridFromPreset creates a grid with the provided configuration.
func GridFromPreset(size usize, values [][]uint64) Grid {
	// if values were provided and are correctly sized
	if values != nil &&
		len(values) == int(size) &&
		len(values[0]) == int(size) {
		return Grid{Size: size, Tiles: toTiles(values)}
	}

	// if values were not provided or are incorrect - create a fresh grid
	grid := make([][]Tile, size)
	for i := range size {
		grid[i] = make([]Tile, size)
	}

	return Grid{Size: size, Tiles: grid}
}

// NewGrid creates an empty grid and spawns two random tiles in it.
func NewGrid(size usize) Grid {
	grid := make([][]Tile, size)
	for i := range size {
		grid[i] = make([]Tile, size)
	}

	g := Grid{Size: size, Tiles: grid}
	g.SpawnTile(DefaultDistribution)
	g.SpawnTile(DefaultDistribution)

	return g
}

// MakeMove performs a full move cycle in the following order:
//  1. Shifts all tiles in the provided direction;
//  2. Combines adjacent tiles with equal values;
//  3. Shifts all tiles in the provided direction;
//  4. Spawns a new random tile, using [Grid.SpawnTile] with [DefaultDistribution].
//
// It returns the score, obtained by this move
func (g *Grid) MakeMove(direction Direction) (score uint64) {
	score = g.Shift(direction)
	g.SpawnTile(DefaultDistribution)
	return
}

// Shift performs a shift in the provided direction.
//
// It returns the score gained from this action.
func (g *Grid) Shift(direction Direction) (score uint64) {
	switch direction {
	case Up:
		return g.shiftUp()
	case Left:
		return g.shiftLeft()
	case Down:
		return g.shiftDown()
	case Right:
		return g.shiftRight()
	}
	return 0
}

// SpawnTile spawns a new Tile on the field.
//
// Tile value is chosen based on the provided distribution.
// In case the distribution isn't provided the default distribution is used (10% for 4, 90% for 2).
//
// Tile position is chosen uniformly and randomly from the list of empty Tiles.
// If the grid is full the spawn does not occur.
func (g *Grid) SpawnTile(distribution []uint64) {
	if distribution == nil {
		distribution = DefaultDistribution
	}

	voids := [][]usize{}
	for i := range g.Size {
		for j := range g.Size {
			if g.Tiles[i][j].IsEmpty() {
				voids = append(voids, []usize{i, j})
			}
		}
	}
	if len(voids) == 0 {
		return
	}
	void := voids[rand.Intn(len(voids))]
	value := distribution[rand.Intn(len(distribution))]

	tile := &g.Tiles[void[0]][void[1]]
	tile.Value, tile.NewlySpawned = value, true
}

// shiftLeft simulates the change to state, made by swiping left.
func (g *Grid) shiftLeft() (score uint64) {
	g.compress()
	score = g.bind()
	g.compress()
	return
}

// shiftRight simulates the change to state, made by swiping right.
func (g *Grid) shiftRight() (score uint64) {
	g.reverse()
	score = g.shiftLeft()
	g.reverse()
	return
}

// shiftUp simulates the change to state, made by swiping up.
func (g *Grid) shiftUp() (score uint64) {
	g.transpose()
	score = g.shiftLeft()
	g.transpose()
	return
}

// shiftDown simulates the change to state, made by swiping down.
func (g *Grid) shiftDown() (score uint64) {
	g.transpose()
	score = g.shiftRight()
	g.transpose()
	return
}

// bind consumes every two horizontally adjacent tiles with equal values in the grid,
// spawning a new tile inplace of the left one.
// The new tile has the combined value of these tiles.
//
// Created tiles cannot be bound in this iteration again
func (g *Grid) bind() (score uint64) {
	for i := range g.Size {
		for j := range g.Size - 1 {
			cur, next := g.Tiles[i][j], g.Tiles[i][j+1]
			if cur == next {
				score += uint64(next.Value * 2)
				g.Tiles[i][j].Value = next.Value * 2
				g.Tiles[i][j+1].makeVoid()
				j++
			}
		}
	}
	return
}

// Alignes all tiles in the grid along the left border.
func (g *Grid) compress() {
	for i := range g.Size {
		var writeIdx usize = 0
		for j := range g.Size {
			g.Tiles[i][j].NewlySpawned = false
			if g.Tiles[i][j].IsEmpty() {
				continue
			}
			if writeIdx != j {
				g.Tiles[i][writeIdx] = g.Tiles[i][j]
				g.Tiles[i][j].makeVoid()
			}
			writeIdx++
		}
	}
}

// Transposes the grid.
func (g *Grid) transpose() {
	for i := range g.Size {
		for j := i + 1; j < g.Size; j++ {
			g.Tiles[i][j], g.Tiles[j][i] = g.Tiles[j][i], g.Tiles[i][j]
		}
	}
}

// Reverses the order of tiles in each row.
func (g *Grid) reverse() {
	for i := range g.Size {
		slices.Reverse(g.Tiles[i])
	}
}
