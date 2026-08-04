package core

// func randomGrid(size usize) Grid {
// 	grid := make([][]Cell, size)
// 	for i := range size {
// 		grid[i] = make([]Cell, size)
// 	}

// 	for i := range grid {
// 		for j := range grid[i] {
// 			rnd := (Cell)(rand.Uint64N(LIMIT))
// 			var powerOfTwo Cell = 1
// 			for powerOfTwo < rnd {
// 				powerOfTwo *= 2
// 			}
// 			grid[i][j] = powerOfTwo
// 		}
// 	}
// 	return Grid{Size: size, Cells: grid}
// }

// var testGrid = NewGrid(4, [][]Cell{
// 	{0, 2, 2, 8},
// 	{0, 0, 2, 4},
// 	{2, 4, 16, 8},
// 	{0, 2, 0, 32},
// })
