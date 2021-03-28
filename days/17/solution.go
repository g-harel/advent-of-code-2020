package solution

import (
	"fmt"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type Coord struct {
	x int
	y int
	z int
}

func NewCoord(str string) Coord {
	parts := lib.ParseInts(strings.Split(str, ","))
	return Coord{parts[0], parts[1], parts[2]}
}

// Produce string representaton of coordinates.
// Will never produce same string from different coordinates.
func (c *Coord) String() string {
	return fmt.Sprintf("%v,%v,%v", c.x, c.y, c.z)
}

//

type Grid struct {
	// Only true coordinates should be populated in the grid.
	// x -> y -> z -> value
	values map[string]bool
}

func NewGrid() Grid {
	return Grid{
		values: map[string]bool{},
	}
}

func (g *Grid) GetCount() int {
	return len(g.values)
}

func (g *Grid) SetMarked(c Coord) {
	g.values[c.String()] = true
}

func (g *Grid) GetMarked(c Coord) bool {
	return g.values[c.String()]
}

// Return list of all marked coordinates.
func (g *Grid) GetAllMarked() []Coord {
	cs := []Coord{}
	for v := range g.values {
		cs = append(cs, NewCoord(v))
	}
	return cs
}

//

func iterate(prev Grid) Grid {
	// Compute list of candidates that could be marked in the next iteration.
	// This list contains unique neighbors of all currently marked coordinates.
	candidates := []Coord{}
	wasIterated := NewGrid()
	for _, marked := range prev.GetAllMarked() {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			for yOffset := -1; yOffset <= 1; yOffset++ {
				for zOffset := -1; zOffset <= 1; zOffset++ {
					neighbor := Coord{
						x: marked.x + xOffset,
						y: marked.y + yOffset,
						z: marked.z + zOffset,
					}
					if wasIterated.GetMarked(neighbor) {
						continue
					}
					wasIterated.SetMarked(neighbor)
					candidates = append(candidates, neighbor)
				}
			}
		}
	}

	// Iterate coordinates in next grid.
	next := NewGrid()
	for _, candidate := range candidates {
		neighbors := 0
		for xOffset := -1; xOffset <= 1; xOffset++ {
			for yOffset := -1; yOffset <= 1; yOffset++ {
				for zOffset := -1; zOffset <= 1; zOffset++ {
					if xOffset == 0 && yOffset == 0 && zOffset == 0 {
						continue
					}
					neighbor := Coord{
						x: candidate.x + xOffset,
						y: candidate.y + yOffset,
						z: candidate.z + zOffset,
					}
					if prev.GetMarked(neighbor) {
						neighbors++
					}
				}
			}
		}
		if prev.GetMarked(candidate) {
			if neighbors == 2 || neighbors == 3 {
				next.SetMarked(candidate)
			}
		} else {
			if neighbors == 3 {
				next.SetMarked(candidate)
			}
		}
	}

	return next
}

//

func Part1() int {
	lines := lib.ReadLines("input.txt")

	grid := NewGrid()
	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				grid.SetMarked(Coord{j, i, 0})
			}
		}
	}

	for i := 0; i < 6; i++ {
		grid = iterate(grid)
	}

	return grid.GetCount()
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
