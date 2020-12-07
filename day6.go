package aoc2020

import "strings"

func Day6Part1() int {
	groups := SplitGroups(ReadLines("day6.input.txt"))

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

func Day6Part2() int {
	groups := SplitGroups(ReadLines("day6.input.txt"))

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
