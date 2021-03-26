package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type validRange struct {
	start int
	end   int
}

func parseInput(lines []string) (ranges map[string][]validRange, yourTicket []int, nearbyTickets [][]int) {
	ranges = map[string][]validRange{}
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
	return ranges, yourTicket, nearbyTickets
}

func Part1() int {
	ranges, _, nearbyTickets := parseInput(lib.ReadLines("input.txt"))

	errors := 0
	for _, nearbyTicket := range nearbyTickets {
		for _, value := range nearbyTicket {
			inRange := false
			for _, validRanges := range ranges {
				if inRange {
					continue
				}
				for _, vr := range validRanges {
					if value >= vr.start && value <= vr.end {
						inRange = true
						continue
					}
				}
			}
			if !inRange {
				errors += value
			}
		}
	}

	return errors
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
