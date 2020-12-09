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

	childToParentsMap := map[string][]string{}
	for _, parent := range bags {
		for _, child := range parent.childNames {
			childToParentsMap[child] = append(childToParentsMap[child], parent.name)
		}
	}

	canBeParent := map[string]bool{}
	var collectPotentialParents func(string)
	collectPotentialParents = func(child string) {
		for _, parent := range childToParentsMap[child] {
			canBeParent[parent] = true
			collectPotentialParents(parent)
		}
	}
	collectPotentialParents("shiny gold")

	return len(canBeParent)
}

func Day7Part2() int {
	bags := day7ParseBags(ReadLines("day7.input.txt"))

	parentToChildNameMap := map[string][]string{}
	parentToChildCountMap := map[string][]int{}
	for _, parent := range bags {
		parentToChildNameMap[parent.name] = parent.childNames
		parentToChildCountMap[parent.name] = parent.childCounts
	}

	var countChildren func(string) int
	countChildren = func(parent string) int {
		total := 1
		for i, child := range parentToChildNameMap[parent] {
			total += countChildren(child) * parentToChildCountMap[parent][i]
		}
		return total
	}

	return countChildren("shiny gold") - 1
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
