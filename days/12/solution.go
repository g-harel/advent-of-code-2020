package solution

import (
	"math"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Part1() int {
	lines := lib.ReadLines("input.txt")

	x, y := 0, 0
	direction := 0 // angle in degrees where zero is due east
	for _, line := range lines {
		action := line[0:1]
		value := lib.ParseInt(line[1:])
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
			y += int(math.Sin(lib.Rad(direction))) * value
			x += int(math.Cos(lib.Rad(direction))) * value
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func Part2() int {
	return -1
}
