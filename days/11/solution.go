package solution

import "github.com/g-harel/advent-of-code-2020/lib"

func Part1() int {
	lines := lib.ReadLines("input.txt")

	prev := lines
	for {
		prev = lines
		lines = runRound(lines)
		if boardEqual(prev, lines) {
			return countOccupied(lines)
		}
	}
}

func Part2() int {
	return -1
}

func runRound(lines []string) []string {
	updated := []string{}
	for i := 0; i < len(lines); i++ {
		line := ""
		for j := 0; j < len(lines[i]); j++ {
			neighbors := countNeighbors(lines, j, i)
			if lines[i][j] == "L"[0] && neighbors == 0 {
				line += "#"
			} else if lines[i][j] == "#"[0] && neighbors >= 4 {
				line += "L"
			} else {
				line += string(lines[i][j])
			}
		}
		updated = append(updated, line)
	}
	return updated
}

func countNeighbors(lines []string, x, y int) int {
	isOccupied := func(x, y int) int {
		if y < 0 || x < 0 || y >= len(lines) || x >= len(lines[y]) {
			return 0
		}
		if lines[y][x] == "#"[0] {
			return 1
		}
		return 0
	}
	return isOccupied(x-1, y-1) + isOccupied(x, y-1) + isOccupied(x+1, y-1) +
		isOccupied(x-1, y) + isOccupied(x+1, y) +
		isOccupied(x-1, y+1) + isOccupied(x, y+1) + isOccupied(x+1, y+1)
}

func boardEqual(a, b []string) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func countOccupied(lines []string) int {
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "#"[0] {
				count++
			}
		}
	}
	return count
}
