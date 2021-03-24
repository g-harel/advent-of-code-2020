package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Part1() int {
	nums := lib.ParseInts(strings.Split(lib.ReadLines("input.txt")[0], ","))
	return iter(nums, 2020)
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}

func iter(nums []int, n int) int {
	lastSpoken := map[int]int{}

	for i, num := range nums {
		lastSpoken[num] = i
	}

	prev := nums[len(nums)-1]
	for i := len(nums); i <= n-1; i++ {
		if _, ok := lastSpoken[prev]; !ok {
			lastSpoken[prev] = i - 1
			prev = 0
			continue
		}
		originalPrev := prev
		prev = i - lastSpoken[prev] - 1
		lastSpoken[originalPrev] = i - 1
	}

	return prev
}
