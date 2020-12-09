package aoc2020

import (
	"testing"
)

func AssertCorrect(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Fatalf("incorrect: expected %v but got %v", expected, actual)
	}
}

func Test(t *testing.T) {
	AssertCorrect(t, Day1Part1(), 712075)
	AssertCorrect(t, Day1Part2(), 145245270)

	AssertCorrect(t, Day2Part1(), 422)
	AssertCorrect(t, Day2Part2(), 451)

	AssertCorrect(t, Day3Part1(), 145)
	AssertCorrect(t, Day3Part2(), 3424528800)

	AssertCorrect(t, Day4Part1(), 213)
	AssertCorrect(t, Day4Part2(), 147)

	AssertCorrect(t, Day5Part1(), 801)
	AssertCorrect(t, Day5Part2(), 597)

	AssertCorrect(t, Day6Part1(), 6782)
	AssertCorrect(t, Day6Part2(), 3596)

	AssertCorrect(t, Day7Part1(), 151)
	AssertCorrect(t, Day7Part2(), 41559)
}
