package str

import "fmt"

type islandCounter struct {
	sand    map[string]bool
	islands int
}

func numIslands(grid [][]byte) int {
	counter := islandCounter{
		sand: make(map[string]bool),
	}
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			coord := fmt.Sprintf("%d,%d", x, y)
			fmt.Println(coord)
			if _, ok := counter.sand[coord]; ok {
				continue
			}

			if grid[x][y] == 49 {
				counter.addIsland()
				counter.mapIsland(grid, x, y)
				fmt.Println("count", counter.islands)
			}
		}
	}
	return counter.islands
}

func (i islandCounter) mapIsland(grid [][]byte, x, y int) {
	coord := fmt.Sprintf("%d,%d", x, y)
	if grid[x][y] == 48 || i.sand[coord] {
		return
	}

	i.sand[coord] = true

	if y > 0 {
		i.mapIsland(grid, x, y-1)
	}
	if x < len(grid)-1 {
		i.mapIsland(grid, x+1, y)
	}
	if y < len(grid[x])-1 {
		i.mapIsland(grid, x, y+1)
	}
	if x > 0 {
		i.mapIsland(grid, x-1, y)
	}
}

func (i *islandCounter) addIsland() {
	i.islands++
}
