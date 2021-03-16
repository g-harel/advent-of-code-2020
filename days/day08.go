package days

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

func Day08Part1() int {
	_, acc := day08Run(lib.ReadLines("day08.input.txt"))
	return acc
}

func Day08Part2() int {
	lines := lib.ReadLines("day08.input.txt")

	for i, line := range lines {
		if strings.Contains(line, "nop") {
			lines[i] = strings.ReplaceAll(line, "nop", "jmp")
			loop, acc := day08Run(lines)
			if !loop {
				return acc
			}
		}
		if strings.Contains(line, "jmp") {
			lines[i] = strings.ReplaceAll(line, "jmp", "nop")
			loop, acc := day08Run(lines)
			if !loop {
				return acc
			}
		}
		lines[i] = line
	}

	return -1
}

func day08ParseInstruction(str string) (jmp bool, arg int) {
	parts := strings.Split(str, " ")
	operation := parts[0]
	argument := lib.ParseSignedInt(parts[1])
	if operation == "nop" {
		operation = "jmp"
		argument = 1
	}
	return operation == "jmp", argument
}

func day08Run(instructions []string) (loop bool, acc int) {
	seenInstructions := map[int]bool{}
	accumulator := 0
	currentInstruction := 0
	for {
		if currentInstruction >= len(instructions) {
			return false, accumulator
		}
		if seenInstructions[currentInstruction] {
			return true, accumulator
		}
		seenInstructions[currentInstruction] = true

		jmp, arg := day08ParseInstruction(instructions[currentInstruction])
		if jmp {
			currentInstruction += arg
			continue
		}
		accumulator += arg
		currentInstruction++
	}
}
