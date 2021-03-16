package days

import (
	"testing"
)

func AssertCorrect(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Fatalf("incorrect: expected %v but got %v", expected, actual)
	}
}

func Test(t *testing.T) {
	AssertCorrect(t, Day05Part1(), 801)
	AssertCorrect(t, Day05Part2(), 597)

	AssertCorrect(t, Day06Part1(), 6782)
	AssertCorrect(t, Day06Part2(), 3596)

	AssertCorrect(t, Day07Part1(), 151)
	AssertCorrect(t, Day07Part2(), 41559)

	AssertCorrect(t, Day08Part1(), 1801)
	AssertCorrect(t, Day08Part2(), 2060)

	AssertCorrect(t, Day09Part1(), 1212510616)
	AssertCorrect(t, Day09Part2(), 171265123)

	AssertCorrect(t, Day10Part1(), 2775)
	AssertCorrect(t, Day10Part2(), -1)

	AssertCorrect(t, Day11Part1(), 2299)
	AssertCorrect(t, Day11Part2(), -1)

	AssertCorrect(t, Day12Part1(), 1482)
	AssertCorrect(t, Day12Part2(), -1)
}
