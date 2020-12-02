package aoc2020

func Day1Part1() int {
	return MapMult(FindNumsThatSum(2020, 2, MapInt(ReadLines("day1.input.txt"))))
}

func Day1Part2() int {
	return MapMult(FindNumsThatSum(2020, 3, MapInt(ReadLines("day1.input.txt"))))
}
