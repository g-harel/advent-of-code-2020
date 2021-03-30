package solution

import (
	"strconv"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type Coord struct {
	values []int
}

func NewCoord(c ...int) Coord {
	return Coord{
		values: c,
	}
}

func ParseCoord(str string) Coord {
	return Coord{values: lib.ParseInts(strings.Split(str, ","))}
}

// Produce string representaton of coordinates.
// Will never generate same string from different coordinates.
func (c *Coord) String() string {
	strs := []string{}
	for _, value := range c.values {
		strs = append(strs, strconv.Itoa(value))
	}
	return strings.Join(strs, ",")
}

// List of neighbor coordinates (at most one away in any direction).
// Also includes self coordinates.
func (c *Coord) Neighbors() []Coord {
	combinations := lib.Combinations(len(c.values), []int{-1, 0, 1})
	neighbors := []Coord{}
	for _, combination := range combinations {
		neighbor := Coord{values: []int{}}
		for i, offset := range combination {
			neighbor.values = append(neighbor.values, c.values[i]+offset)
		}
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
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
		cs = append(cs, ParseCoord(v))
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
		for _, neighbor := range marked.Neighbors() {
			if wasIterated.GetMarked(neighbor) {
				continue
			}
			wasIterated.SetMarked(neighbor)
			candidates = append(candidates, neighbor)
		}
	}

	// Iterate coordinates in next grid.
	next := NewGrid()
	for _, candidate := range candidates {
		neighbors := 0
		for _, neighbor := range candidate.Neighbors() {
			if prev.GetMarked(neighbor) {
				neighbors++
			}
		}
		if prev.GetMarked(candidate) {
			neighbors--
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
				grid.SetMarked(NewCoord(j, i, 0))
			}
		}
	}

	for i := 0; i < 6; i++ {
		grid = iterate(grid)
	}

	return grid.GetCount()
}

func Part2() int {
	lines := lib.ReadLines("input.txt")

	grid := NewGrid()
	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				grid.SetMarked(NewCoord(j, i, 0, 0))
			}
		}
	}

	for i := 0; i < 6; i++ {
		grid = iterate(grid)
	}

	return grid.GetCount()
}
