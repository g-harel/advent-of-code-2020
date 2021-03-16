if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
	exit 1
fi

mkdir "days/$1"

cat <<EOF >> "days/$1/input.txt"
EOF

cat <<EOF >> "days/$1/solution.go"
package solution

import "lib"

func Part1() int {
	lib.ReadLines("input.txt")
	return -1
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
EOF

cat <<EOF >> "days/$1/solution_test.go"
package solution_test

import (
	"lib"
	"testing"

	solution "github.com/g-harel/advent-of-code-2020/days/$1"
)

func Test(t *testing.T) {
	lib.AssertCorrect(t, solution.Part1(), -1)
	lib.AssertCorrect(t, solution.Part2(), -1)
}
EOF
