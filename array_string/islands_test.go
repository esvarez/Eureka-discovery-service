package str

import (
	"fmt"
	"testing"
)

func TestIslandCounter(t *testing.T) {
	grid := [][]byte{
		{49, 49, 49, 49, 48},
		{49, 49, 48, 49, 48},
		{49, 49, 48, 48, 48},
		{48, 48, 48, 48, 48},
	}

	out := numIslands(grid)

	fmt.Println(out)

}
