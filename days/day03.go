package days

import "lib"

func Day03Part1() int {
	return day03CalcCollisions(lib.ReadLines("day03.input.txt"), 3, 1)
}

func Day03Part2() int {
	lines := lib.ReadLines("day03.input.txt")
	return day03CalcCollisions(lines, 1, 1) *
		day03CalcCollisions(lines, 3, 1) *
		day03CalcCollisions(lines, 5, 1) *
		day03CalcCollisions(lines, 7, 1) *
		day03CalcCollisions(lines, 1, 2)
}

func day03CalcCollisions(lines []string, incX, incY int) int {
	count := 0
	for i := incY; i < len(lines); i += incY {
		if day03IsTree(lines[i], i/incY*incX) {
			count++
		}
	}
	return count
}

func day03IsTree(line string, pos int) bool {
	return line[pos%len(line)] == []byte("#")[0]
}
