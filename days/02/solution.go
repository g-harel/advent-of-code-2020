package solution

import (
	"strings"

	"lib"
)

func Part1() int {
	return applyPolicy(lib.ReadLines("input.txt"), rentalPolicy)
}

func Part2() int {
	return applyPolicy(lib.ReadLines("input.txt"), tobogganPolicy)
}

func applyPolicy(lines []string, policy func(a, b int, letter, password string) bool) int {
	count := 0
	for _, line := range lines {
		spaceSplit := strings.Split(line, " ")
		rangeSplit := strings.Split(spaceSplit[0], "-")

		a := lib.ParseInt(rangeSplit[0])
		b := lib.ParseInt(rangeSplit[1])
		letter := strings.ReplaceAll(spaceSplit[1], ":", "")
		password := spaceSplit[2]

		if policy(a, b, letter, password) {
			count++
		}
	}
	return count
}

func rentalPolicy(a, b int, letter, password string) bool {
	count := 0
	for _, char := range password {
		if char == []rune(letter)[0] {
			count++
		}
	}
	return count >= a && count <= b
}

func tobogganPolicy(a, b int, letter, password string) bool {
	letterByte := []byte(letter)[0]
	return (password[a-1] == letterByte) != (password[b-1] == letterByte)
}
