package aoc2020

import (
	"regexp"
	"strings"
)

type day4Checker func(string) bool

func Day4Part1() int {
	return day4CountCheckers(ReadLines("day4.input.txt"), map[string][]day4Checker{
		"byr": {day4CheckIsPresent},
		"iyr": {day4CheckIsPresent},
		"eyr": {day4CheckIsPresent},
		"hgt": {day4CheckIsPresent},
		"hcl": {day4CheckIsPresent},
		"ecl": {day4CheckIsPresent},
		"pid": {day4CheckIsPresent},
	})
}

func Day4Part2() int {
	return day4CountCheckers(ReadLines("day4.input.txt"), map[string][]day4Checker{
		"byr": {day4GenCheckRegex(`^\d{4}$`), day4GenCheckRange(1920, 2002)},
		"iyr": {day4GenCheckRegex(`^\d{4}$`), day4GenCheckRange(2010, 2020)},
		"eyr": {day4GenCheckRegex(`^\d{4}$`), day4GenCheckRange(2020, 2030)},
		"hgt": {day4GenCheckRegex(`^\d{2,3}\w{2}$`), day4CheckHgt},
		"hcl": {day4GenCheckRegex(`^#[0-9a-f]{6}$`)},
		"ecl": {day4GenCheckRegex(`^(amb|blu|brn|gry|grn|hzl|oth)$`)},
		"pid": {day4GenCheckRegex(`^\d{9}$`)},
	})
}

func day4CountCheckers(lines []string, checkers map[string][]day4Checker) int {
	count := 0
	for _, passport := range day4ExtractPassports(lines) {
		if day4ValidatePassport(passport, checkers) {
			count++
		}
	}
	return count
}

func day4CheckIsPresent(val string) bool {
	return val != ""
}

func day4CheckHgt(val string) bool {
	if strings.HasSuffix(val, "cm") {
		num := ParseInt(strings.TrimSuffix(val, "cm"))
		return num >= 150 && num <= 193
	}
	if strings.HasSuffix(val, "in") {
		num := ParseInt(strings.TrimSuffix(val, "in"))
		return num >= 59 && num <= 76
	}
	return false
}

func day4GenCheckRange(min, max int) day4Checker {
	return func(val string) bool {
		num := ParseInt(val)
		return num >= min && num <= max
	}
}

func day4GenCheckRegex(pattern string) day4Checker {
	return func(val string) bool {
		re := regexp.MustCompile(pattern)
		return re.MatchString(val)
	}
}

func day4ExtractPassports(lines []string) []map[string]string {
	passports := SplitGroups(lines)

	parsedPassports := []map[string]string{}
	for _, passport := range passports {
		parsedPassport := map[string]string{}
		for _, passportLine := range passport {
			parts := strings.Split(strings.TrimSpace(passportLine), " ")
			for _, part := range parts {
				entries := strings.Split(part, ":")
				parsedPassport[entries[0]] = entries[1]
			}
		}
		parsedPassports = append(parsedPassports, parsedPassport)
	}
	return parsedPassports
}

func day4ValidatePassport(passport map[string]string, check map[string][]day4Checker) bool {
	for field, checkers := range check {
		for _, checker := range checkers {
			if !checker(passport[field]) {
				return false
			}
		}
	}
	return true
}
