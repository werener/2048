package core

import (
	"errors"
	"fmt"
)

type usize uint16
type number uint64

var (
	IndexOutOfRange = errors.New("Index out of range")
)

const LIMIT = 2048

type Grid struct {
	Size  usize
	Cells [][]Cell
}

func NewGrid(size usize, values [][]number) Grid {
	grid := make([][]Cell, size)
	for i := range size {
		grid[i] = make([]Cell, size)
	}

	if values != nil {
		for i := range size {
			for j := range size {
				cell := &grid[i][j]

				cell.Value = values[i][j]
				if cell.Value == 0 {
					cell.Empty = true
				}
			}
		}
	}
	return Grid{Size: size, Cells: grid}
}

func (grid *Grid) row(i usize) ([]Cell, error) {
	if i >= grid.Size {
		return []Cell{}, IndexOutOfRange
	}
	return grid.Cells[i], nil
}

func (grid *Grid) col(j usize) ([]Cell, error) {
	if j >= grid.Size {
		return []Cell{}, IndexOutOfRange
	}

	col := make([]Cell, grid.Size)
	for i := range grid.Cells {
		col[i] = grid.Cells[i][j]
	}
	return col, nil
}

func (grid *Grid) Print() {
	for _, row := range grid.Cells {
		for _, item := range row {
			fmt.Print(item.Value, " ")
		}
		println()
	}
}

type Cell struct {
	Empty bool
	Value number
}

func NewCell() Cell {
	return Cell{Empty: true}
}

func mapToValues(cells []Cell) []number {
	values := make([]number, len(cells))
	for i := range cells {
		values[i] = cells[i].Value
	}
	return values
}
