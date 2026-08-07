package models

import (
	"math/rand"
)

// Grid represents the aggregation of playing tiles.
type Grid struct {
	Size  usize    // size of the grid. Grid is always (Size x Size)
	Tiles [][]Tile // 2d array of all tiles
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

// NewGrid creates an empty grid.
func NewGrid(size usize) Grid {
	grid := make([][]Tile, size)
	for i := range size {
		grid[i] = make([]Tile, size)
	}

	g := Grid{Size: size, Tiles: grid}

	return g
}

// SpawnTile spawns a new Tile on the field.
//
// Tile value is chosen based on the provided distribution.
// In case the distribution isn't provided the [DefaultDistribution] is used.
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

	g.Tiles[void[0]][void[1]] = randomTile(distribution)
}

// Clone returns a deep copy of the grid
func (g *Grid) Clone() Grid {
	tiles := make([][]Tile, g.Size)
	for i := range g.Size {
		tiles[i] = make([]Tile, g.Size)
		copy(tiles[i], g.Tiles[i])
	}
	return Grid{
		Size:  g.Size,
		Tiles: tiles,
	}
}
