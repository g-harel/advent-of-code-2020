package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type bagDefinition struct {
	name        string
	childNames  []string
	childCounts []int
}

func Part1() int {
	bags := parseBags(lib.ReadLines("input.txt"))

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

func Part2() int {
	bags := parseBags(lib.ReadLines("input.txt"))

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

func parseBags(lines []string) []bagDefinition {
	bags := []bagDefinition{}
	for _, line := range lines {
		line = strings.ReplaceAll(line, " bags contain ", ":")
		line = strings.ReplaceAll(line, " bags.", "")
		line = strings.ReplaceAll(line, " bag.", "")
		line = strings.ReplaceAll(line, " bags, ", ":")
		line = strings.ReplaceAll(line, " bag, ", ":")
		line = strings.ReplaceAll(line, ":no other", "")

		parts := strings.Split(line, ":")
		bag := bagDefinition{
			name: parts[0],
		}
		for _, child := range parts[1:] {
			childSplit := strings.SplitN(child, " ", 2)
			bag.childNames = append(bag.childNames, childSplit[1])
			bag.childCounts = append(bag.childCounts, lib.ParseInt(childSplit[0]))
		}
		bags = append(bags, bag)
	}
	return bags
}
