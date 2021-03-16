package solution

import "lib"

func Part1() int {
	return calcCollisions(lib.ReadLines("input.txt"), 3, 1)
}

func Part2() int {
	lines := lib.ReadLines("input.txt")
	return calcCollisions(lines, 1, 1) *
		calcCollisions(lines, 3, 1) *
		calcCollisions(lines, 5, 1) *
		calcCollisions(lines, 7, 1) *
		calcCollisions(lines, 1, 2)
}

func calcCollisions(lines []string, incX, incY int) int {
	count := 0
	for i := incY; i < len(lines); i += incY {
		if isTree(lines[i], i/incY*incX) {
			count++
		}
	}
	return count
}

func isTree(line string, pos int) bool {
	return line[pos%len(line)] == []byte("#")[0]
}
