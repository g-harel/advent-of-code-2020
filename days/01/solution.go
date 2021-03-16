package solution

import "lib"

func Part1() int {
	return lib.MapMult(lib.FindNumsThatSum(2020, 2, lib.MapInt(lib.ReadLines("input.txt"))))
}

func Part2() int {
	return lib.MapMult(lib.FindNumsThatSum(2020, 3, lib.MapInt(lib.ReadLines("input.txt"))))
}
