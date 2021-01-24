if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
	exit 1
fi

cat <<EOF >> "day$1.input.txt"
EOF

cat <<EOF >> "day$1.go"
package aoc2020

func Day$1Part1() int {
	return -1
}

func Day$1Part2() int {
	return -1
}
EOF
