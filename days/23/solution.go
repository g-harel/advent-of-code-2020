package solution

import (
	"fmt"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func WrapGet(nums []int, i int) int {
	return nums[i%len(nums)]
}

func WrapRange(nums []int, count int) []int {
	ii := []int{}
	i := 0
	for len(ii) < count {
		ii = append(ii, i%len(nums))
		i++
	}
	return ii
}

func GetSmaller(nums []int, i int) int {
	val := nums[i] - 1
	min, max := lib.MinMax(nums)
	if val == min-1 {
		val = max
	}
	return lib.IndexOf(nums, val)
}

func Part1() int {
	cups := lib.ParseInts(strings.Split(lib.ReadLines("input.txt")[0], ""))

	for _, i := range WrapRange(cups, 100) {
		destinationIndex := GetSmaller(cups, i)
		for destinationIndex > i && destinationIndex < i+3 {
			destinationIndex = GetSmaller(cups, destinationIndex)
		}

		// TODO scoot over the three next nums.
		fmt.Println(cups, WrapGet(cups, i), WrapGet(cups, destinationIndex))
	}

	return 0
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
