package solution_test

import (
	"testing"

	solution "github.com/g-harel/advent-of-code-2020/days/12"
	"github.com/g-harel/advent-of-code-2020/lib"
)

func Test(t *testing.T) {
	lib.AssertCorrect(t, solution.Part1(), 1482)
	lib.AssertCorrect(t, solution.Part2(), -1)
}
