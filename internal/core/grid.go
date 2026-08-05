package core

import (
	"math/rand"
	"slices"
)

type usize uint16
type Cell uint64
type Direction int8

const (
	Up Direction = iota
	Left
	Down
	Right
)
const EMPTY Cell = 0

var DefaultDistribution = []Cell{4, 2, 2, 2, 2, 2, 2, 2, 2, 2}

// Grid represents the aggregation of playing blocks.
type Grid struct {
	Size  usize
	Cells [][]Cell
}

// GridFromPreset creates a grid with the provided configuration.
func GridFromPreset(size usize, values [][]Cell) Grid {
	// if values were provided and are correctly sized
	if values != nil &&
		len(values) == int(size) &&
		len(values[0]) == int(size) {
		return Grid{Size: size, Cells: values}
	}

	// if values were not provided or are incorrect - create a fresh grid
	grid := make([][]Cell, size)
	for i := range size {
		grid[i] = make([]Cell, size)
	}

	return Grid{Size: size, Cells: grid}
}

// NewGrid creates an empty grid and spawns two random cells in it.
func NewGrid(size usize) Grid {
	grid := make([][]Cell, size)
	for i := range size {
		grid[i] = make([]Cell, size)
	}

	g := Grid{Size: size, Cells: grid}
	g.SpawnCell(DefaultDistribution)
	g.SpawnCell(DefaultDistribution)

	return g
}

// MakeMove performs a full move cycle in the following order:
//  1. Shifts all cells in the provided direction;
//  2. Combines adjacent cells with equal values;
//  3. Shifts all cells in the provided direction;
//  4. Spawns a new random cell, using [Grid.SpawnCell] with [DefaultDistribution].
//
// It returns the score, obtained by this move
func (g *Grid) MakeMove(direction Direction) (score uint64) {
	score = g.Shift(direction)
	g.SpawnCell(DefaultDistribution)
	return
}

// Shift performs a shift in the provided direction.
//
// It returns the score gained from this action.
func (g *Grid) Shift(direction Direction) (score uint64) {
	switch direction {
	case Up:
		score = g.shiftUp()
	case Left:
		score = g.shiftLeft()
	case Down:
		score = g.shiftDown()
	case Right:
		score = g.shiftRight()
	}
	return 0
}

// SpawnCell spawns a new Cell on the field.
//
// Cell value is chosen based on the provided distribution.
// In case the distribution isn't provided the default distribution is used (10% for 4, 90% for 2).
//
// Cell position is chosen uniformly and randomly from the list of Cells.
// If the grid is full the spawn does not occur.
func (g *Grid) SpawnCell(distribution []Cell) {
	if distribution == nil {
		distribution = DefaultDistribution
	}

	voids := [][]usize{}
	for i := range g.Size {
		for j := range g.Size {
			if g.Cells[i][j] == EMPTY {
				voids = append(voids, []usize{i, j})
			}
		}
	}
	if len(voids) == 0 {
		return
	}
	void := voids[rand.Intn(len(voids))]
	value := distribution[rand.Intn(len(distribution))]

	g.Cells[void[0]][void[1]] = value
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

// bind consumes every two horizontally adjacent cells with equal values in the grid,
// spawning a new cell inplace of the left one.
// The new cell has the combined value of these cells.
//
// Created cells cannot be bound in this iteration again
func (g *Grid) bind() (score uint64) {
	for i := range g.Size {
		for j := range g.Size - 1 {
			cur, next := g.Cells[i][j], g.Cells[i][j+1]
			if cur == next {
				score = uint64(next * 2)
				g.Cells[i][j] = next * 2
				g.Cells[i][j+1] = EMPTY
				j++
			}
		}
	}
	return
}

// Alignes all cells in the grid along the left border.
func (g *Grid) compress() {
	for i := range g.Size {
		var writeIdx usize = 0
		for j := range g.Size {
			if g.Cells[i][j] == EMPTY {
				continue
			}
			if writeIdx != j {
				g.Cells[i][writeIdx] = g.Cells[i][j]
				g.Cells[i][j] = EMPTY
			}
			writeIdx++
		}
	}
}

// Transposes the grid.
func (g *Grid) transpose() {
	for i := range g.Size {
		for j := i + 1; j < g.Size; j++ {
			g.Cells[i][j], g.Cells[j][i] = g.Cells[j][i], g.Cells[i][j]
		}
	}
}

// Reverses the order of cells in each row.
func (g *Grid) reverse() {
	for i := range g.Size {
		slices.Reverse(g.Cells[i])
	}
}
