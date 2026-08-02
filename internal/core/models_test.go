package core

import (
	"math/rand/v2"
	"testing"
)

func randomGrid(size usize) Grid {
	grid := make([][]Cell, size)
	for i := range size {
		grid[i] = make([]Cell, size)
	}

	for i := range grid {
		for j := range grid[i] {
			rnd := (number)(rand.Uint64N(LIMIT))
			var powerOfTwo number = 1
			for powerOfTwo < rnd {
				powerOfTwo *= 2
			}
			grid[i][j].Value = powerOfTwo
		}
	}
	return Grid{Size: size, Cells: grid}
}

var testGrid = NewGrid(4, [][]number{
	{0, 2, 2, 8},
	{0, 0, 2, 4},
	{2, 4, 16, 8},
	{0, 2, 0, 32},
})

func TestRow(t *testing.T) {
	tests := []struct {
		name           string
		i              usize
		expectedResult []number
		expectedError  error
	}{
		// {"fetch row", 1, []number{0, 0, 2, 4}, nil},
		{"fetch row", 4, []number{}, IndexOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := testGrid.row(tt.i)

			if err != tt.expectedError {
				t.Errorf("\nExpected error '%v' to appear\nGot %v", tt.expectedError, err)
			}

			for i := range result {
				if result[i].Value != tt.expectedResult[i] {
					t.Errorf("\nExpected %v\nGot %v (index %d doesn't match)", mapToValues(result), tt.expectedResult, i)
				}
			}
		})
	}
}
