package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type LoopSlice struct {
	nums []int
}

func (l *LoopSlice) Min() int {
	min := l.nums[0]
	for _, n := range l.nums {
		if n < min {
			min = n
		}
	}
	return min
}

func (l *LoopSlice) Max() int {
	max := 0
	for _, n := range l.nums {
		if n > max {
			max = n
		}
	}
	return max
}

func (l *LoopSlice) Loop(n int) []int {
	ii := []int{}
	i := 0
	for len(ii) < n {
		ii = append(ii, i%len(l.nums))
		i++
	}
	return ii
}

func Part1() int {
	cups := lib.ParseInts(strings.Split(lib.ReadLines("input.txt")[0], ""))
	max := len(cups)

	for i := 0; i < 100; i++ {
		current := cups[i]%max - 1
		three := cups[i+1 : i+4] // TODO can go out of bounds.

		target := current - 1
		if target == 0 {
			target = max
		}

		for lib.IndexOf(three, target) > 0 {
			current--
			if current == 0 {
				current = max
			}
		}

		// destination := lib.IndexOf(cups, target)
	}

	return -1
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
