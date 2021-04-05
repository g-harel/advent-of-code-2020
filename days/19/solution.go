package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

// Consumes n items in str starting at index i.
// Returns 0 when nothing is consumed.
type Consumer = func(i int, str string) (n int)

func ByteConsumer(match byte) Consumer {
	return func(i int, str string) int {
		if str[i] == match {
			return 1
		}
		return 0
	}
}

func SuccessiveConsumer(consumers ...Consumer) Consumer {
	return func(i int, str string) int {
		nTotal := 0
		for _, consumer := range consumers {
			n := consumer(nTotal+i, str)
			if n == 0 {
				return 0
			}
			nTotal += n
		}
		return nTotal
	}
}

func EitherConsumer(consumers ...Consumer) Consumer {
	return func(i int, str string) int {
		for _, consumer := range consumers {
			n := consumer(i, str)
			if n != 0 {
				return n
			}
		}
		return 0
	}
}

func DeferredConsumer(fn func() Consumer) Consumer {
	return func(i int, str string) int {
		return fn()(i, str)
	}
}

func ParseConsumer(lines []string) Consumer {
	rules := map[string]Consumer{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		id := parts[0]
		rule := parts[1]

		if strings.HasPrefix(rule, "\"") {
			rules[id] = ByteConsumer(rule[1])
			continue
		}

		eitherConditions := strings.Split(rule, " | ")
		eitherConsumers := []Consumer{}
		for _, eitherCondition := range eitherConditions {
			eitherIDs := strings.Split(eitherCondition, " ")
			successiveConsumers := []Consumer{}
			for _, eitherID := range eitherIDs {
				localEitherID := eitherID
				successiveConsumers = append(
					successiveConsumers,
					DeferredConsumer(func() Consumer {
						return rules[localEitherID]
					}),
				)
			}
			eitherConsumers = append(eitherConsumers, SuccessiveConsumer(successiveConsumers...))
		}
		rules[id] = EitherConsumer(eitherConsumers...)
	}

	return rules["0"]
}

func Part1() int {
	groups := lib.SplitGroups(lib.ReadLines("input.txt"))

	matches := 0
	for _, message := range groups[1] {
		if ParseConsumer(groups[0])(0, message) == len(message) {
			matches++
		}
	}

	return matches
}

func Part2() int {
	lines := lib.ReadLines("input.txt")
	for i, line := range lines {
		if line == "8: 42" {
			lines[i] = "8: 42 | 42 8"
		}
		if line == "11: 42 31" {
			lines[i] = "11: 42 31 | 42 11 31"
		}
	}

	groups := lib.SplitGroups(lines)

	matches := 0
	for _, message := range groups[1] {
		if ParseConsumer(groups[0])(0, message) == len(message) {
			matches++
		}
	}

	return -1
}
