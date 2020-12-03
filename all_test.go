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
}
