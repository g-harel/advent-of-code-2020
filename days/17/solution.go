package solution

import "github.com/g-harel/advent-of-code-2020/lib"

type Grid struct {
	// Only active coordinates should be populated in the grid.
	// x -> y -> z -> active
	active map[int]map[int]map[int]bool
}

func (g *Grid) SetActive(x, y, z int) {
	if g.active == nil {
		g.active = map[int]map[int]map[int]bool{}
	}
	if _, ok := g.active[x]; !ok {
		g.active[x] = map[int]map[int]bool{}
	}
	if _, ok := g.active[x][y]; !ok {
		g.active[x][y] = map[int]bool{}
	}
	g.active[x][y][z] = true
}

func (g *Grid) Count() int {
	sum := 0
	for _, xs := range g.active {
		for _, ys := range xs {
			sum += len(ys)
		}
	}
	return sum
}

func Part1() int {
	lines := lib.ReadLines("input.txt")

	grid := Grid{}
	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				grid.SetActive(j, i, 0)
			}
		}
	}

	for i := 0; i < 0; i++ {
		grid = iterate(grid)
	}

	return grid.Count()
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}

func iterate(grid Grid) Grid {
	next := Grid{}
	return next
}
