package solution

import (
	"fmt"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type WrapSlice struct {
	nums []int
}

func (w *WrapSlice) Get(i int) int {
	return w.nums[i%len(w.nums)]
}

func (w *WrapSlice) IndexOf(n int) int {
	return lib.IndexOf(w.nums, n)
}

func (w *WrapSlice) Min() int {
	min := w.nums[0]
	for _, n := range w.nums {
		if n < min {
			min = n
		}
	}
	return min
}

func (w *WrapSlice) Max() int {
	max := 0
	for _, n := range w.nums {
		if n > max {
			max = n
		}
	}
	return max
}

func (w *WrapSlice) Range(n int) []int {
	ii := []int{}
	i := 0
	for len(ii) < n {
		ii = append(ii, i%len(w.nums))
		i++
	}
	return ii
}

func GetSmaller(w *WrapSlice, i int) int {
	val := w.Get(i) - 1
	if val == w.Min()-1 {
		val = w.Max()
	}
	return w.IndexOf(val)
}

func Part1() int {
	cups := &WrapSlice{lib.ParseInts(strings.Split(lib.ReadLines("input.txt")[0], ""))}

	for _, i := range cups.Range(100) {
		destinationIndex := GetSmaller(cups, i)
		for destinationIndex > i && destinationIndex < i+3 {
			destinationIndex = GetSmaller(cups, destinationIndex)
		}

		// TODO scoot over the three next nums.
		fmt.Println(cups, cups.Get(i), cups.Get(destinationIndex))
	}

	return 0
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
