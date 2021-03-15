package aoc2020

import "math"

func Day12Part1() int {
	lines := ReadLines("day12.input.txt")

	x, y := 0, 0
	direction := 0 // angle in degrees where zero is due east
	for _, line := range lines {
		action := line[0:1]
		value := ParseInt(line[1:])
		switch action {
		case "N":
			y += value
		case "S":
			y -= value
		case "E":
			x += value
		case "W":
			x -= value
		case "L":
			direction += value
		case "R":
			direction -= value
		case "F":
			y += int(math.Sin(Rad(direction))) * value
			x += int(math.Cos(Rad(direction))) * value
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func Day12Part2() int {
	return -1
}
