package aoc2020

import (
	"strings"
)

func Day2Part1() int {
	return day2ApplyPolicy(ReadLines("day2.input.txt"), day2RentalPolicy)
}

func Day2Part2() int {
	return day2ApplyPolicy(ReadLines("day2.input.txt"), day2TobogganPolicy)
}

func day2ApplyPolicy(lines []string, policy func(a, b int, letter, password string) bool) int {
	count := 0
	for _, line := range lines {
		spaceSplit := strings.Split(line, " ")
		rangeSplit := strings.Split(spaceSplit[0], "-")

		a := ParseInt(rangeSplit[0])
		b := ParseInt(rangeSplit[1])
		letter := strings.ReplaceAll(spaceSplit[1], ":", "")
		password := spaceSplit[2]

		if policy(a, b, letter, password) {
			count++
		}
	}
	return count
}

func day2RentalPolicy(a, b int, letter, password string) bool {
	count := 0
	for _, char := range password {
		if char == []rune(letter)[0] {
			count++
		}
	}
	return count >= a && count <= b
}

func day2TobogganPolicy(a, b int, letter, password string) bool {
	letterByte := []byte(letter)[0]
	return (password[a-1] == letterByte) != (password[b-1] == letterByte)
}
