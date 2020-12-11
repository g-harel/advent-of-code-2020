package aoc2020

func Day01Part1() int {
	return MapMult(FindNumsThatSum(2020, 2, MapInt(ReadLines("day01.input.txt"))))
}

func Day01Part2() int {
	return MapMult(FindNumsThatSum(2020, 3, MapInt(ReadLines("day01.input.txt"))))
}
