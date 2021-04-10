package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Edges(lines []string) []string {
	left := ""
	right := ""
	for _, line := range lines {
		left += string(line[0])
		right += string(line[len(line)-1])
	}
	return []string{lines[0], lines[len(lines)-1], left, right}
}

func Part1() int {
	groups := lib.SplitGroups(lib.ReadLines("input.txt"))

	tiles := map[int][]string{}
	for _, group := range groups {
		id := strings.TrimSuffix(strings.TrimPrefix(group[0], "Tile "), ":")
		tiles[lib.ParseInt(id)] = group[1:]
	}

	edges := map[int][]string{}
	for id, lines := range tiles {
		edges[id] = Edges(lines)
	}

	matchedEdges := map[int]int{}
	for id, groupEdges := range edges {
		for _, edge := range groupEdges {
			for compareID, compareGroupEdges := range edges {
				for _, compareEdge := range compareGroupEdges {
					if id == compareID {
						continue
					}
					if edge == compareEdge || edge == lib.Reverse(compareEdge) {
						matchedEdges[id]++
					}
				}
			}
		}
	}

	product := 1
	for id, count := range matchedEdges {
		if count == 2 {
			product *= id
		}
	}

	return product
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
