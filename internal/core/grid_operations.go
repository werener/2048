package core

import "slices"

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
