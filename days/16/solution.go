package solution

import (
	"fmt"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type validRange struct {
	start int
	end   int
}

func Part1() int {
	lines := lib.ReadLines("input.txt")

	ranges := map[string][]validRange{}
	yourTicket := []int{}
	nearbyTickets := [][]int{}
	section := 0
	for _, line := range lines {
		if line == "your ticket:" || line == "nearby tickets:" {
			continue
		}
		if line == "" {
			section++
			continue
		}
		if section == 0 {
			firstSplit := strings.Split(line, ": ") // class: 1-3 or 5-7
			field := firstSplit[0]
			validRanges := strings.Split(firstSplit[1], " or ")
			for _, rawRange := range validRanges {
				rangeParts := strings.Split(rawRange, "-")
				ranges[field] = append(ranges[field], validRange{
					start: lib.ParseInt(rangeParts[0]),
					end:   lib.ParseInt(rangeParts[1]),
				})
			}
			continue
		}
		if section == 1 {
			yourTicket = lib.ParseInts(strings.Split(line, ","))
			continue
		}
		nearbyTickets = append(nearbyTickets, lib.ParseInts(strings.Split(line, ",")))
	}

	fmt.Println(ranges, yourTicket, nearbyTickets)

	return 1
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
