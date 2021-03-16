package days

import "lib"

func Day01Part1() int {
	return lib.MapMult(lib.FindNumsThatSum(2020, 2, lib.MapInt(lib.ReadLines("day01.input.txt"))))
}

func Day01Part2() int {
	return lib.MapMult(lib.FindNumsThatSum(2020, 3, lib.MapInt(lib.ReadLines("day01.input.txt"))))
}
