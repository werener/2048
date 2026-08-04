package main

import (
	"fmt"

	_ "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core"
)

func main() {
	grid := core.NewGrid(4, [][]core.Cell{
		{2, 0, 2, 8},
		{2, 0, 2, 4},
		{2, 4, 16, 8},
		{4, 2, 4, 4},
	})

	fmt.Println(grid.Repr())

	grid.ShiftLeft()
	fmt.Printf("Shift Left\n%s\n", grid.Repr())

	grid.ShiftRight()
	fmt.Printf("Shift Right\n%s\n", grid.Repr())

	grid.ShiftUp()
	fmt.Printf("Shift Up\n%s\n", grid.Repr())

	grid.ShiftDown()
	fmt.Printf("Shift Down\n%s\n", grid.Repr())
}
