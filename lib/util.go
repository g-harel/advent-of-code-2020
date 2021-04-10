package lib

import (
	"math"
	"strings"
)

func Rad(deg int) float64 {
	return float64(deg) * (math.Pi / 180.0)
}

func Multiply(nums []int) int {
	product := 1
	for _, num := range nums {
		product *= num
	}
	return product
}

func FindNumsThatSum(target, count int, nums []int) []int {
	// Find pair that equals target value.
	if count == 2 {
		seen := map[int]bool{}
		for _, num := range nums {
			missing := target - num
			if seen[missing] {
				return []int{missing, num}
			}
			seen[num] = true
		}
		return []int{}
	}

	for i, num := range nums {
		// Create slice with current num missing.
		remaining := make([]int, len(nums))
		copy(remaining, nums)
		remaining = append(remaining[:i], remaining[i+1:]...)

		// Look in remaining for nums that would equal missing.
		missing := target - num
		result := FindNumsThatSum(missing, count-1, remaining)
		if len(result) > 0 {
			return append(result, num)
		}
	}

	return []int{}
}

func SplitGroups(lines []string) [][]string {
	groups := [][]string{}
	currentGroup := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			groups = append(groups, currentGroup)
			currentGroup = []string{}
			continue
		}
		currentGroup = append(currentGroup, line)
	}
	groups = append(groups, currentGroup)
	return groups
}

func Combinations(positions int, options []int) [][]int {
	combinations := [][]int{{}}
	for i := 0; i < positions; i++ {
		newCombinations := [][]int{}
		for _, option := range options {
			for _, prevCombination := range combinations {
				newCombination := make([]int, len(prevCombination))
				copy(newCombination, prevCombination)
				newCombinations = append(newCombinations, append(newCombination, option))
			}
		}
		combinations = newCombinations
	}
	return combinations
}

func Reverse(str string) string {
	result := ""
	for _, r := range str {
		result = string(r) + result
	}
	return result
}

func Common(a, b []string) []string {
	m := map[string]bool{}
	for _, v := range a {
		m[v] = true
	}
	common := []string{}
	for _, v := range b {
		if m[v] {
			common = append(common, v)
		}
	}
	return common
}

func Remove(strs []string, mask ...string) []string {
	out := []string{}
	for _, s := range strs {
		match := false
		for _, m := range mask {
			if s == m {
				match = true
				break
			}
		}
		if !match {
			out = append(out, s)
		}
	}
	return out
}

func Unique(strs []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, str := range strs {
		if seen[str] {
			continue
		}
		seen[str] = true
		out = append(out, str)
	}
	return out
}
