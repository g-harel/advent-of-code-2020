package solution

import "github.com/g-harel/advent-of-code-2020/lib"

func Part1() int {
	return lib.Multiply(lib.FindNumsThatSum(2020, 2, lib.ParseInts(lib.ReadLines("input.txt"))))
}

func Part2() int {
	return lib.Multiply(lib.FindNumsThatSum(2020, 3, lib.ParseInts(lib.ReadLines("input.txt"))))
}
