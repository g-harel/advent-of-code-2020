package lib

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

func ParseSignedInt(str string) int {
	if strings.HasPrefix(str, "+") {
		return ParseInt(str[1:])
	} else if strings.HasPrefix(str, "-") {
		return -ParseInt(str[1:])
	} else {
		return ParseInt(str)
	}
}

func ParseBinary(str string) int {
	num, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(fmt.Errorf("parse binary string: %v", err))
	}
	return int(num)
}

func ParseInts(strs []string) []int {
	var nums []int
	for _, str := range strs {
		nums = append(nums, ParseInt(str))
	}
	return nums
}
