package core

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type usize uint16
type Cell uint64

const EMPTY Cell = 0

var (
	IndexOutOfRange = errors.New("Index out of range")
)

// Grid represents the aggregation of playing blocks
type Grid struct {
	Size  usize
	Cells [][]Cell
}

// NewGrid returns a new grid of provided size
func NewGrid(size usize, values [][]Cell) Grid {
	grid := make([][]Cell, size)
	for i := range size {
		grid[i] = make([]Cell, size)
	}

	if values != nil {
		for i := range size {
			for j := range size {
				grid[i][j] = values[i][j]
			}
		}
	}
	return Grid{Size: size, Cells: grid}
}

// ShiftLeft simulates the change to state, made by swiping left
func (g *Grid) ShiftLeft() {
	g.compress()
	g.bind()
	g.compress()
}

// ShiftRight simulates the change to state, made by swiping left
func (g *Grid) ShiftRight() {
	g.reverse()
	g.ShiftLeft()
	g.reverse()
}

// ShiftUp simulates the change to state, made by swiping left
func (g *Grid) ShiftUp() {
	g.transpose()
	g.ShiftLeft()
	g.transpose()
}

// ShiftDown simulates the change to state, made by swiping left
func (g *Grid) ShiftDown() {
	g.transpose()
	g.ShiftRight()
	g.transpose()
}

// Repr returns a string representation of the grid
func (g *Grid) Repr() string {
	var buf strings.Builder

	for i := range g.Size {
		for j := range g.Size {
			fmt.Fprintf(&buf, "%d ", g.Cells[i][j])
		}
		fmt.Fprintln(&buf)
	}

	return buf.String()
}

// Binds adjacent cells in the grid together
func (g *Grid) bind() {
	for i := range g.Size {
		for j := range g.Size - 1 {
			cur, next := g.Cells[i][j], g.Cells[i][j+1]
			if cur == next {
				g.Cells[i][j] = next * 2
				g.Cells[i][j+1] = EMPTY
				j++
			}
		}
	}
}

// Alignes all cells in the grid along the left border
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

// Transposes the grid
func (g *Grid) transpose() {
	for i := range g.Size {
		for j := i + 1; j < g.Size; j++ {
			g.Cells[i][j], g.Cells[j][i] = g.Cells[j][i], g.Cells[i][j]
		}
	}
}

// Reverses the order of cells in each row
func (g *Grid) reverse() {
	for i := range g.Size {
		slices.Reverse(g.Cells[i])
	}
}
