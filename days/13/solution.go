package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Part1() int {
	lines := lib.ReadLines("input.txt")

	time := lib.ParseInt(lines[0])
	buses := []int{}
	for _, b := range strings.Split(lines[1], ",") {
		if b != "x" {
			buses = append(buses, lib.ParseInt(b))
		}
	}

	earliestBus := -1
	shortestWait := -1
	for _, bus := range buses {
		currentWait := bus - time%bus
		if currentWait < shortestWait || shortestWait == -1 {
			earliestBus = bus
			shortestWait = currentWait
		}
	}

	return earliestBus * shortestWait
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
