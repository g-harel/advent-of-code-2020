package aoc2020

import "strings"

func Day8Part1() int {
	lines := ReadLines("day8.input.txt")

	seenInstructions := map[int]bool{}
	accumulator := 0
	currentInstruction := 0
	for {
		if seenInstructions[currentInstruction] {
			return accumulator
		}
		seenInstructions[currentInstruction] = true

		jmp, arg := day8ParseInstruction(lines[currentInstruction])
		if jmp {
			currentInstruction += arg
			continue
		}
		accumulator += arg
		currentInstruction++
	}
}

func Day8Part2() int {
	return -1
}

func day8ParseInstruction(str string) (jmp bool, arg int) {
	parts := strings.Split(str, " ")
	operation := parts[0]
	argument := ParseSignedInt(parts[1])
	if operation == "nop" {
		operation = "jmp"
		argument = 1
	}
	return operation == "jmp", argument
}
