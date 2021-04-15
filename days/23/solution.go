package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func WrapGet(nums []int, i int) int {
	return nums[lib.Mod(i, len(nums))]
}

func WrapIter(nums []int, count int) []int {
	indecies := []int{}
	i := 0
	for len(indecies) < count {
		indecies = append(indecies, lib.Mod(i, len(nums)))
		i++
	}
	return indecies
}

func WrapRange(nums []int, a, b int) []int {
	r := []int{}
	for i := a; i < b; i++ {
		ii := lib.Mod(i, len(nums))
		r = append(r, nums[ii])
	}
	return r
}

func Part1() int {
	cups := lib.ParseInts(strings.Split(lib.ReadLines("input.txt")[0], ""))

	for _, i := range WrapIter(cups, 100) {
		min, max := lib.MinMax(cups)
		rotated := lib.Rotate(cups, -i)[4:]

		destination := cups[i]
		destinationIndex := i
		for {
			destination--
			if destination < min {
				destination = max
			}
			destinationIndex = lib.IndexOf(rotated, destination)
			if destinationIndex >= 0 {
				break
			}
		}

		newCups := lib.Concat(
			[]int{cups[i]},
			rotated[:destinationIndex+1],
			WrapRange(cups, i+1, i+4),
			rotated[destinationIndex+1:])

		cups = lib.Rotate(newCups, i)
	}

	// Format answer.
	return lib.ParseInt(strings.Join(lib.ParseStrings(lib.Rotate(cups, -lib.IndexOf(cups, 1))[1:]), ""))
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
