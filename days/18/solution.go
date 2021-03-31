package solution

import "github.com/g-harel/advent-of-code-2020/lib"

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

func (a *Atom) Evaluate() int {
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

func Part1() int {
	lib.ReadLines("input.txt")

	one := 1
	two := 2
	three := 3

	// 1+2*3
	first := Atom{
		expression: &Expression{
			operator: "*",
			left: Atom{
				expression: &Expression{
					operator: "+",
					left: Atom{
						value: &one,
					},
					right: Atom{
						value: &two,
					},
				},
			},
			right: Atom{
				value: &three,
			},
		},
	}

	// 1+(2*3)
	second := Atom{
		expression: &Expression{
			operator: "+",
			left: Atom{
				value: &one,
			},
			right: Atom{
				expression: &Expression{
					operator: "*",
					left: Atom{
						value: &two,
					},
					right: Atom{
						value: &three,
					},
				},
			},
		},
	}

	return first.Evaluate() + second.Evaluate()
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
