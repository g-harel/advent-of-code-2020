package days

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Day06Part1() int {
	groups := lib.SplitGroups(lib.ReadLines("day06.input.txt"))

	total := 0
	for _, group := range groups {
		answers := map[string]bool{}
		for _, person := range group {
			for _, letter := range strings.Split(person, "") {
				answers[letter] = true
			}
		}
		total += len(answers)
	}
	return total
}

func Day06Part2() int {
	groups := lib.SplitGroups(lib.ReadLines("day06.input.txt"))

	total := 0
	for _, group := range groups {
		answers := map[string]int{}
		for _, person := range group {
			for _, letter := range strings.Split(person, "") {
				answers[letter]++
			}
		}
		target := len(group)
		for _, answers := range answers {
			if answers == target {
				total++
			}
		}
	}
	return total
}
