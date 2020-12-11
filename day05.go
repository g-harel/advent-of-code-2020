package aoc2020

import (
	"sort"
	"strings"
)

func Day05Part1() int {
	max := 0
	for _, id := range day05DecodeAll(ReadLines("day05.input.txt")) {
		if id > max {
			max = id
		}
	}
	return max
}

func Day05Part2() int {
	ids := day05DecodeAll(ReadLines("day05.input.txt"))
	sort.Ints(ids)

	offset := ids[0]
	for i, id := range ids {
		if id != i+offset {
			return id - 1
		}
	}
	return 0
}

func day05DecodeAll(lines []string) []int {
	ids := []int{}
	for _, line := range lines {
		row, col := day05Decode(line)
		ids = append(ids, row*8+col)
	}
	return ids
}

func day05Decode(seat string) (int, int) {
	row := ParseBinary(strings.ReplaceAll(strings.ReplaceAll(seat[:7], "B", "1"), "F", "0"))
	column := ParseBinary(strings.ReplaceAll(strings.ReplaceAll(seat[7:], "R", "1"), "L", "0"))
	return row, column
}
