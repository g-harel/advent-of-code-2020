package solution_test

import (
	"lib"
	"testing"

	solution "github.com/g-harel/advent-of-code-2020/days/04"
)

func Test(t *testing.T) {
	lib.AssertCorrect(t, solution.Part1(), 213)
	lib.AssertCorrect(t, solution.Part2(), 147)
}
