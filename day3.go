package aoc2020

func Day3Part1() int {
	return day3CalcCollisions(ReadLines("day3.input.txt"), 3, 1)
}

func Day3Part2() int {
	lines := ReadLines("day3.input.txt")
	return day3CalcCollisions(lines, 1, 1) *
		day3CalcCollisions(lines, 3, 1) *
		day3CalcCollisions(lines, 5, 1) *
		day3CalcCollisions(lines, 7, 1) *
		day3CalcCollisions(lines, 1, 2)
}

func day3CalcCollisions(lines []string, incX, incY int) int {
	count := 0
	for i := incY; i < len(lines); i += incY {
		if day3IsTree(lines[i], i/incY*incX) {
			count++
		}
	}
	return count
}

func day3IsTree(line string, pos int) bool {
	return line[pos%len(line)] == []byte("#")[0]
}
