package core

import "fmt"

const emptyValue uint64 = 0

type Tile struct {
	Value        uint64
	NewlySpawned bool
}

func newTile(value uint64) Tile {
	return Tile{Value: value, NewlySpawned: false}
}
func toTiles(values [][]uint64) [][]Tile {
	n, m := len(values), len(values[0])

	grid := make([][]Tile, n)
	for i := range n {
		grid[i] = make([]Tile, m)
		for j := range m {
			grid[i][j] = newTile(values[i][j])
		}
	}
	return grid
}

func (t Tile) IsEmpty() bool {
	return t.Value == emptyValue
}

func (t Tile) Repr() string {
	if t.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%d", t.Value)
}

func (t *Tile) makeVoid() {
	t.Value = emptyValue
}
