package solution_test

import (
	"testing"

	solution "github.com/g-harel/advent-of-code-2020/days/16"
	"github.com/g-harel/advent-of-code-2020/lib"
)

func Test(t *testing.T) {
	lib.AssertCorrect(t, solution.Part1(), 23054)
	lib.AssertCorrect(t, solution.Part2(), -1)
}
