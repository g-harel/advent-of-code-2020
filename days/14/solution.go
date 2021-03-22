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

	mask := strings.ReplaceAll(lines[0], "mask = ", "")
	ops := []writeOp{}
	for _, opLine := range lines[1:] {
		s := strings.Split(opLine, " = ")
		addr := strings.ReplaceAll(strings.ReplaceAll(s[0], "mem[", ""), "]", "")
		ops = append(ops, writeOp{
			address: lib.ParseInt(addr),
			value:   lib.ParseInt(s[1]),
		})
	}

	mem := map[int]int{}
	applyMaskedOps(mem, mask, ops)

	sum := 0
	for addr, val := range mem {
		println(addr)
		println(val)
		println("===")
		sum += val
	}

	return sum
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}

func applyMaskedOps(mem map[int]int, mask string, writes []writeOp) {
	oneMask := lib.ParseBinary(strings.ReplaceAll(mask, "X", "0"))
	zeroMask := lib.ParseBinary(strings.ReplaceAll(mask, "X", "1"))
	for _, w := range writes {
		mem[w.address] = (w.value | oneMask) & zeroMask
	}
}
