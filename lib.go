package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AssertCorrect(actual, expected interface{}) {
	if actual != expected {
		panic(fmt.Errorf("incorrect: expected %v but got %v", expected, actual))
	}
}

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open file: %v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(fmt.Errorf("scan file: %v", err))
	}
	return lines
}

func MapInt(strs []string) []int {
	var nums []int
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Errorf("convert int: %v", err))
		}
		nums = append(nums, num)
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
