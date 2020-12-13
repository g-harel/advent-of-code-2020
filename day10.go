package aoc2020

import (
	"sort"
)

func Day10Part1() int {
	nums := MapInt(ReadLines("day10.input.txt"))
	sort.Ints(nums)

	numOneDiff := 0
	numThreeDiff := 1 // last jump
	last := 0         // initial outlet
	for _, num := range nums {
		if num-last == 1 {
			numOneDiff++
		}
		if num-last == 3 {
			numThreeDiff++
		}
		last = num
	}

	return numOneDiff * numThreeDiff
}

func Day10Part2() int {
	return -1
}
