package core

import "slices"

// transformation is a reversable function, that rearranges the [Tile] 2D array in a grid.
//
// transformation(transformation(g)) = g
type transformation func(*Grid)

/*
transformations relates the provided [Direction] of a shift to a slice of [transformation] functions,
which should be consecutively applied before [Grid.bind]
to perform a shift in the provided direction by performing a left shift.

Usage example:

	for _, transformation := range transformations[direction] {
		transformation(grid)
		defer transformation(grid)
	}
	score := grid.bind()
*/
var transformations = map[Direction][]transformation{
	Left:  {},
	Right: {reverse},
	Up:    {transpose},
	Down:  {transpose, reverse},
}

// Shift performs a shift in the provided [Direction].
//
// It returns the score gained from this action.
func (g *Grid) Shift(direction Direction) (score uint64) {
	for _, transformation := range transformations[direction] {
		transformation(g)
		defer transformation(g)
	}
	g.compress()
	defer g.compress()

	return g.bind()
}

// canShift checks if [Grid.Shift] in the provided [Direction] would make any changes to the grid.
//
// Does not mutate the grid. Pass by pointer is for performance reasons.
func (g *Grid) canShift(direction Direction) bool {
	for _, transformation := range transformations[direction] {
		transformation(g)
		defer transformation(g)
	}
	return g.canBind() || g.canCompress()

}

// canBind checks whether [Grid.bind] would change the state of the grid.
//
// Does not mutate the grid. Pass by pointer is for performance reasons.
func (g *Grid) canBind() bool {
	for i := range g.Size {
		var lastValue uint64 = 1
		for j := range g.Size {
			if g.Tiles[i][j].IsEmpty() {
				continue
			}
			if g.Tiles[i][j].Value == lastValue {
				return true
			}
			lastValue = g.Tiles[i][j].Value
		}
	}
	return false
}

// bind consumes every two horizontally adjacent tiles with equal values in the grid,
// spawning a new tile in place of the left one.
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

// canBind checks whether [Grid.compress] would change the state of the grid.
//
// Does not mutate the grid. Pass by pointer is for performance reasons.
func (g *Grid) canCompress() bool {
	for i := range g.Size {
		rowHasVoid := false
		for j := range g.Size {
			if g.Tiles[i][j].IsEmpty() {
				rowHasVoid = true
			} else {
				if rowHasVoid {
					return true
				}
			}
		}
	}
	return false
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
func transpose(g *Grid) {
	for i := range g.Size {
		for j := i + 1; j < g.Size; j++ {
			g.Tiles[i][j], g.Tiles[j][i] = g.Tiles[j][i], g.Tiles[i][j]
		}
	}
}

// Reverses the order of tiles in each row.
func reverse(g *Grid) {
	for i := range g.Size {
		slices.Reverse(g.Tiles[i])
	}
}
