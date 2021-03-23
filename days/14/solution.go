package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type writeOp struct {
	address int
	value   int
}

func Part1() int {
	lines := lib.ReadLines("input.txt")

	mem := map[int]int{}
	mask := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "mask = ") {
			mask = strings.ReplaceAll(line, "mask = ", "")
			continue
		}

		parts := strings.Split(line, " = ")
		addr := strings.ReplaceAll(strings.ReplaceAll(parts[0], "mem[", ""), "]", "")
		mem[lib.ParseInt(addr)] = applyMask(mask, lib.ParseInt(parts[1]))
	}

	sum := 0
	for _, val := range mem {
		sum += val
	}

	return sum
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}

func applyMask(mask string, value int) int {
	oneMask := lib.ParseBinary(strings.ReplaceAll(mask, "X", "0"))
	zeroMask := lib.ParseBinary(strings.ReplaceAll(mask, "X", "1"))
	return (value | oneMask) & zeroMask
}
