package solution

import (
	"fmt"
	"strconv"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type Expression struct {
	operator string
	left     Atom
	right    Atom
}

// One of expression or value (branch vs. leaf).
type Atom struct {
	expression *Expression
	value      *int
}

func (a Atom) Print() string {
	if a.value != nil {
		return strconv.Itoa(*a.value)
	}
	if a.expression == nil {
		return "#ERR-BOTH-NULL"
	}
	return fmt.Sprintf("(%v%v%v)",
		a.expression.left.Print(),
		a.expression.operator,
		a.expression.right.Print())
}

func (a Atom) Evaluate() int {
	if a.value != nil {
		return *a.value
	}
	if a.expression == nil {
		return -1
	}
	left := a.expression.left.Evaluate()
	right := a.expression.right.Evaluate()
	switch a.expression.operator {
	case "+":
		return left + right
	case "*":
		return left * right
	}
	return -1
}

//

// Assumes expression is correct.
func Parse(str string) Atom {
	current := Atom{}
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		switch char {
		case " ":
			continue
		case "(":
			start := i
			nestCount := 1
			for nestCount != 0 {
				i++
				subChar := string(str[i])
				if subChar == ")" {
					nestCount--
				} else if subChar == "(" {
					nestCount++
				}
			}
			subExpression := Parse(str[start+1 : i])
			if current.expression != nil {
				current.expression.right = subExpression
			} else {
				current = subExpression
			}
		case "+", "*":
			current = Atom{
				expression: &Expression{
					operator: char,
					left:     current,
				},
			}
		default:
			value := lib.ParseInt(char)
			if current.expression != nil {
				current.expression.right = Atom{value: &value}
			} else {
				current.value = &value
			}
		}
	}
	return current
}

//

func Part1() int {
	lines := lib.ReadLines("input.txt")

	sum := 0
	for _, line := range lines {
		sum += Parse(line).Evaluate()
	}

	return sum
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
