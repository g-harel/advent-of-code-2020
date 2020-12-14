package aoc2020

func Day11Part1() int {
	lines := ReadLines("day11.input.txt")

	prev := lines
	for {
		prev = lines
		lines = day11RunRound(lines)
		if day11BoardEqual(prev, lines) {
			return day11CountOccupied(lines)
		}
	}
}

func Day11Part2() int {
	return -1
}

func day11RunRound(lines []string) []string {
	updated := []string{}
	for i := 0; i < len(lines); i++ {
		line := ""
		for j := 0; j < len(lines[i]); j++ {
			neighbors := day11CountNeighbors(lines, j, i)
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

func day11CountNeighbors(lines []string, x, y int) int {
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

func day11BoardEqual(a, b []string) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func day11CountOccupied(lines []string) int {
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
