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

func Part1() int {
	groups := lib.SplitGroups(lib.ReadLines("input.txt"))

	rules := map[string]Consumer{}
	for _, line := range groups[0] {
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
				successiveConsumers = append(successiveConsumers, rules[eitherID])
			}
			eitherConsumers = append(eitherConsumers, SuccessiveConsumer(successiveConsumers...))
		}
		rules[id] = EitherConsumer(eitherConsumers...)
	}

	matches := 0
	for _, message := range groups[1] {
		if rules["0"](0, message) == len(message) {
			matches++
		}
	}

	return matches
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
