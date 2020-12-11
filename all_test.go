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
	AssertCorrect(t, Day01Part1(), 712075)
	AssertCorrect(t, Day01Part2(), 145245270)

	AssertCorrect(t, Day02Part1(), 422)
	AssertCorrect(t, Day02Part2(), 451)

	AssertCorrect(t, Day03Part1(), 145)
	AssertCorrect(t, Day03Part2(), 3424528800)

	AssertCorrect(t, Day04Part1(), 213)
	AssertCorrect(t, Day04Part2(), 147)

	AssertCorrect(t, Day05Part1(), 801)
	AssertCorrect(t, Day05Part2(), 597)

	AssertCorrect(t, Day06Part1(), 6782)
	AssertCorrect(t, Day06Part2(), 3596)

	AssertCorrect(t, Day07Part1(), 151)
	AssertCorrect(t, Day07Part2(), 41559)

	AssertCorrect(t, Day08Part1(), 1801)
	AssertCorrect(t, Day08Part2(), 2060)

	AssertCorrect(t, Day09Part1(), 1212510616)
	AssertCorrect(t, Day09Part2(), -1)
}
