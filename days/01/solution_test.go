package solution_test

import (
	"lib"
	"testing"

	solution "github.com/g-harel/advent-of-code-2020/days/01"
)

func Test(t *testing.T) {
	lib.AssertCorrect(t, solution.Part1(), 712075)
	lib.AssertCorrect(t, solution.Part2(), 145245270)
}
