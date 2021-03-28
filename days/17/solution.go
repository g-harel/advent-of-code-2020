package solution

import "github.com/g-harel/advent-of-code-2020/lib"

type Grid struct {
	// Only true coordinates should be populated in the grid.
	// x -> y -> z -> value
	values map[int]map[int]map[int]bool
	count  int
}

func (g *Grid) GetCount() int {
	return g.count
}

func (g *Grid) Mark(x, y, z int) {
	if g.values == nil {
		g.values = map[int]map[int]map[int]bool{}
	}
	if _, ok := g.values[x]; !ok {
		g.values[x] = map[int]map[int]bool{}
	}
	if _, ok := g.values[x][y]; !ok {
		g.values[x][y] = map[int]bool{}
	}
	g.values[x][y][z] = true
	g.count++
}

func (g *Grid) IsMarked(x, y, z int) bool {
	if g.values == nil {
		return false
	}
	if _, ok := g.values[x]; !ok {
		return false
	}
	if _, ok := g.values[x][y]; !ok {
		return false
	}
	return g.values[x][y][z]
}

// Return list of all marked x, y, z coordinates.
func (g *Grid) GetAllMarked() [][]int {
	coordinates := [][]int{}
	if g.values == nil {
		return coordinates
	}
	for x, ys := range g.values {
		for y, zs := range ys {
			for z, _ := range zs {
				coordinates = append(coordinates, []int{x, y, z})
			}
		}
	}
	return coordinates
}

func Part1() int {
	lines := lib.ReadLines("input.txt")

	grid := Grid{}
	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				grid.Mark(j, i, 0)
			}
		}
	}

	for i := 0; i < 0; i++ {
		grid = iterate(grid)
	}

	return grid.GetCount()
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}

func iterate(grid Grid) Grid {
	next := Grid{}
	return next
}
