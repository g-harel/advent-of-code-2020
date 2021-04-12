package solution

import (
	"github.com/g-harel/advent-of-code-2020/lib"
)

func Score(cards []int) int {
	score := 0
	for i := 0; i < len(cards); i++ {
		score += cards[len(cards)-1-i] * (i + 1)
	}
	return score
}

func Part1() int {
	lines := lib.ReadLines("input.txt")
	groups := lib.SplitGroups(lines)

	player1 := lib.ParseInts(groups[0][1:])
	player2 := lib.ParseInts(groups[1][1:])

	round := 0
	for len(player1) > round && len(player2) > round {
		card1 := player1[round]
		card2 := player2[round]

		if card1 > card2 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}

		round++
	}

	if len(player1) > len(player2) {
		return Score(player1[round:])
	} else {
		return Score(player2[round:])
	}
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
