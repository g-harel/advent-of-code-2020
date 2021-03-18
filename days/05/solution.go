package solution

import (
	"sort"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Part1() int {
	max := 0
	for _, id := range decodeAll(lib.ReadLines("input.txt")) {
		if id > max {
			max = id
		}
	}
	return max
}

func Part2() int {
	ids := decodeAll(lib.ReadLines("input.txt"))
	sort.Ints(ids)

	offset := ids[0]
	for i, id := range ids {
		if id != i+offset {
			return id - 1
		}
	}
	return 0
}

func decodeAll(lines []string) []int {
	ids := []int{}
	for _, line := range lines {
		row, col := decode(line)
		ids = append(ids, row*8+col)
	}
	return ids
}

func decode(seat string) (int, int) {
	row := lib.ParseBinary(strings.ReplaceAll(strings.ReplaceAll(seat[:7], "B", "1"), "F", "0"))
	column := lib.ParseBinary(strings.ReplaceAll(strings.ReplaceAll(seat[7:], "R", "1"), "L", "0"))
	return row, column
}
