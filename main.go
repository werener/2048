package main

import (
	"fmt"

	_ "charm.land/bubbletea/v2"
	"github.com/werener/2048.git/internal/core"
)

func main() {
	grid := core.NewGrid(4, nil)
	fmt.Println(grid.Repr())

	grid.ShiftLeft()
	fmt.Println(grid.Repr())

}
