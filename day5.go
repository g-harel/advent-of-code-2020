package aoc2020

import (
	"strings"
)

func Day5Part1() int {
	lines := ReadLines("day5.input.txt")

	max := 0
	for _, line := range lines {
		row, col := day5Decode(line)
		val := row*8 + col
		if val > max {
			max = val
		}
	}
	return max
}

func day5Decode(seat string) (int, int) {
	row := ParseBinary(strings.ReplaceAll(strings.ReplaceAll(seat[:7], "B", "1"), "F", "0"))
	column := ParseBinary(strings.ReplaceAll(strings.ReplaceAll(seat[7:], "R", "1"), "L", "0"))
	return row, column
}

func Day5Part2() int {
	return 0
}
