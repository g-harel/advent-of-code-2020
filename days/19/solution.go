package solution

import (
	"fmt"
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type Consumer struct {
	// Consumes n items in str starting at index i.
	// Returns 0 when nothing is consumed.
	consume func(i int, str string, parents []string) (n int)
	id      string
}

func (c *Consumer) Consume(i int, str string, parents []string) int {
	p := strings.Join(parents, ",")
	if c.id != "" {
		parents = append(parents, c.id)
	}
	n := c.consume(i, str, parents)
	if c.id != "" {
		if n == 0 {
			fmt.Printf("| %vFAIL | %v,%v\n", strings.Repeat(" ", len(str)-2), p, c.id)
		} else {
			fmt.Printf("| %v %v %v | %v,%v\n", str[:i], str[i:i+n], str[i+n:], p, c.id)
		}
	}
	return n
}

func ByteConsumer(match byte) *Consumer {
	return &Consumer{
		func(i int, str string, parents []string) int {
			if str[i] == match {
				return 1
			}
			return 0
		}, "b"}
}

func SuccessiveConsumer(consumers ...*Consumer) *Consumer {
	return &Consumer{
		func(i int, str string, parents []string) int {
			nTotal := 0
			for _, consumer := range consumers {
				if i+nTotal >= len(str) {
					return nTotal
				}
				n := consumer.Consume(nTotal+i, str, parents)
				if n == 0 {
					return 0
				}
				nTotal += n
			}
			return nTotal
		}, "s"}
}

func EitherConsumer(consumers ...*Consumer) *Consumer {
	return &Consumer{
		func(i int, str string, parents []string) int {
			for _, consumer := range consumers {
				n := consumer.Consume(i, str, parents)
				if n != 0 {
					return n
				}
			}
			return 0
		}, "e"}
}

func DeferredConsumer(fn func() *Consumer) *Consumer {
	return &Consumer{
		func(i int, str string, parents []string) int {
			return fn().Consume(i, str, parents)
		}, ""}
}

func ParseConsumer(lines []string) *Consumer {
	rules := map[string]*Consumer{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		id := parts[0]
		rule := parts[1]

		if strings.HasPrefix(rule, "\"") {
			rules[id] = ByteConsumer(rule[1])
			rules[id].id = id
			continue
		}

		eitherConditions := strings.Split(rule, " | ")
		eitherConsumers := []*Consumer{}
		for _, eitherCondition := range eitherConditions {
			eitherIDs := strings.Split(eitherCondition, " ")
			successiveConsumers := []*Consumer{}
			for _, eitherID := range eitherIDs {
				localEitherID := eitherID
				successiveConsumers = append(
					successiveConsumers,
					DeferredConsumer(func() *Consumer {
						return rules[localEitherID]
					}),
				)
			}
			eitherConsumers = append(eitherConsumers, SuccessiveConsumer(successiveConsumers...))
		}
		rules[id] = EitherConsumer(eitherConsumers...)
		rules[id].id = id
	}

	return rules["0"]
}

func Part1() int {
	groups := lib.SplitGroups(lib.ReadLines("input.txt"))

	matches := 0
	for _, message := range groups[1] {
		if ParseConsumer(groups[0]).Consume(0, message, []string{}) == len(message) {
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
		if ParseConsumer(groups[0]).Consume(0, message, []string{}) == len(message) {
			println(message)
			matches++
		}
	}

	return matches
}
