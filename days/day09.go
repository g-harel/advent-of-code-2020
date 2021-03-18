package days

import (
	"sort"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Day09Part1() int {
	nums := lib.ParseInts(lib.ReadLines("day09.input.txt"))
	return nums[day09FindNotSumIndex(nums, 25)]
}

func Day09Part2() int {
	nums := lib.ParseInts(lib.ReadLines("day09.input.txt"))
	target := nums[day09FindNotSumIndex(nums, 25)]

	matchingRange := []int{}
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				matchingRange = nums[i:j]
				break
			}
		}
		if len(matchingRange) > 0 {
			break
		}
	}

	sort.Ints(matchingRange)
	return matchingRange[0] + matchingRange[len(matchingRange)-1]
}

func day09FindNotSumIndex(nums []int, bufferSize int) int {
	for i := bufferSize; i <= len(nums); i++ {
		found := false
		for j := 0; j < bufferSize; j++ {
			for k := j + 1; k < bufferSize; k++ {
				if nums[i] == nums[i-j-1]+nums[i-k-1] {
					found = true
				}
			}
		}
		if !found {
			return i
		}
	}
	return -1
}
