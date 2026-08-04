package core

import (
	"errors"
	"fmt"
	"strings"
)

type usize uint16
type Cell uint64

var (
	IndexOutOfRange = errors.New("Index out of range")
)

const LIMIT = 2048

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

// ShiftLeft simulates the change to state, made by swiping left
func (g *Grid) ShiftLeft() {
	for i := range g.Size {
		row := g.Cells[i]

		prevCellId := 0
		firstEmptySpace := 0

		for j := range g.Size {
			cell := row[j]
			prevCell := row[prevCellId]

			if !isEmpty(prevCell) && prevCell == cell {
				row[firstEmptySpace] = cell * 2

			}
		}
	}
}

// Checks if a cell is empty
func isEmpty(cell Cell) bool {
	return cell == 0
}
