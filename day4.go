package aoc2020

import (
	"strings"
)

func Day4Part1() int {
	lines := ReadLines("day4.input.txt")

	passports := []string{}
	currentPassport := ""
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			passports = append(passports, currentPassport)
			currentPassport = ""
			continue
		}
		currentPassport += " " + line
	}
	passports = append(passports, currentPassport)

	count := 0
	for _, passport := range passports {
		if day4IsPassportValid(passport) {
			count++
		}
	}
	return count
}

func day4IsPassportValid(passport string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	parts := strings.Split(passport, " ")
	for _, requiredField := range requiredFields {
		hasField := false
		for _, part := range parts {
			if strings.HasPrefix(part, requiredField) {
				hasField = true
			}
		}
		if !hasField {
			return false
		}
	}
	return true
}
