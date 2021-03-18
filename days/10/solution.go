package solution

import (
	"sort"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Part1() int {
	nums := lib.ParseInts(lib.ReadLines("input.txt"))
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

func Part2() int {
	return -1
}
