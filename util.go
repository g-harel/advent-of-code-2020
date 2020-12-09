package aoc2020

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open file: %v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "//ENDTEST" {
			break
		}
		lines = append(lines, text)
	}
	if scanner.Err() != nil {
		panic(fmt.Errorf("scan file: %v", err))
	}
	return lines
}

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Errorf("convert int: %v", err))
	}
	return num
}

func ParseBinary(str string) int {
	num, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(fmt.Errorf("parse binary string: %v", err))
	}
	return int(num)
}

func MapInt(strs []string) []int {
	var nums []int
	for _, str := range strs {
		nums = append(nums, ParseInt(str))
	}
	return nums
}

func MapMult(nums []int) int {
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
