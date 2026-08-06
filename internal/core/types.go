package core

/*
	TYPES SECTION
*/

// transformation is a reversable function, that rearranges the [Tile] 2D array in a grid.
//
// transformation(transformation(g)) = g
type transformation func(*Grid)

type usize uint16 // usize shows that a value is used for sizes

type Direction int8 // Direction enum type representing one of the four cardinal directions
const (
	Up Direction = iota
	Left
	Down
	Right
)

/*
	GLOBAL VALUES SECTION
*/

// emptyTileValue determines what value represents that a [Tile] is empty
const emptyTileValue uint64 = 0

// DefaultDistribution is an array that represents
// the distribution of new tile spawns in a normal game
var DefaultDistribution = []uint64{4, 2, 2, 2, 2, 2, 2, 2, 2, 2}

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
