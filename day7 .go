package aoc2020

import (
	"strings"
)

type day7BagDefinition struct {
	name        string
	childNames  []string
	childCounts []int
}

func Day7Part1() int {
	bags := day7ParseBags(ReadLines("day7.input.txt"))

	// Map of content->parent(s)
	invertedParentMap := map[string][]string{}
	for _, parent := range bags {
		for _, child := range parent.childNames {
			invertedParentMap[child] = append(invertedParentMap[child], parent.name)
		}
	}

	canBeParent := map[string]bool{}
	var collectPotentialParents func(string)
	collectPotentialParents = func(child string) {
		for _, parent := range invertedParentMap[child] {
			canBeParent[parent] = true
			collectPotentialParents(parent)
		}
	}
	collectPotentialParents("shiny gold")

	return len(canBeParent)
}

func Day7Part2() int {
	return -1
}

func day7ParseBags(lines []string) []day7BagDefinition {
	bags := []day7BagDefinition{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, " bags contain ", ":")
		line = strings.ReplaceAll(line, " bags.", "")
		line = strings.ReplaceAll(line, " bag.", "")
		line = strings.ReplaceAll(line, " bags, ", ":")
		line = strings.ReplaceAll(line, " bag, ", ":")
		line = strings.ReplaceAll(line, ":no other", "")

		parts := strings.Split(line, ":")
		bag := day7BagDefinition{
			name: parts[0],
		}
		for _, child := range parts[1:] {
			childSplit := strings.SplitN(child, " ", 2)
			bag.childNames = append(bag.childNames, childSplit[1])
			bag.childCounts = append(bag.childCounts, ParseInt(childSplit[0]))
		}
		bags = append(bags, bag)
	}
	return bags
}
