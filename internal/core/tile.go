package core

import (
	"fmt"
	"math/rand"
)

const emptyValue uint64 = 0

var DefaultDistribution = []uint64{4, 2, 2, 2, 2, 2, 2, 2, 2, 2} // Array, representing the default distribution of new tile spawns

// Tile represents one Tile in a 2048 grid.
type Tile struct {
	Value        uint64 // value on the tile
	NewlySpawned bool   // whether this tile was spawned before the current move
}

// IsEmpty checks if the tile contains its zero-value
func (t Tile) IsEmpty() bool {
	return t.Value == emptyValue
}

// Repr returns a string representation of the tile
func (t Tile) Repr() string {
	if t.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%d", t.Value)
}

// randomTile returns a random Tile, according to the provided distribution.
//
// If no distribution was provided, [DefaultDistribution] is used.
func randomTile(distribution []uint64) Tile {
	return Tile{
		Value:        distribution[rand.Intn(len(distribution))],
		NewlySpawned: true,
	}
}

// makeVoid sets the tile to its zero-value
func (t *Tile) makeVoid() {
	t.Value = emptyValue
}

// newTile creates a new tile from a provided value.
// This Tile will be considered [Tile.NewlySpawned]
func newTile(value uint64) Tile {
	return Tile{Value: value, NewlySpawned: true}
}

// toTiles converts a 2d array of values into a 2d array of Tiles
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
