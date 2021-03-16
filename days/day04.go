package days

import (
	"regexp"
	"strings"

	"lib"
)

type day04Checker func(string) bool

func Day04Part1() int {
	return day04CountCheckers(lib.ReadLines("day04.input.txt"), map[string][]day04Checker{
		"byr": {day04CheckIsPresent},
		"iyr": {day04CheckIsPresent},
		"eyr": {day04CheckIsPresent},
		"hgt": {day04CheckIsPresent},
		"hcl": {day04CheckIsPresent},
		"ecl": {day04CheckIsPresent},
		"pid": {day04CheckIsPresent},
	})
}

func Day04Part2() int {
	return day04CountCheckers(lib.ReadLines("day04.input.txt"), map[string][]day04Checker{
		"byr": {day04GenCheckRegex(`^\d{4}$`), day04GenCheckRange(1920, 2002)},
		"iyr": {day04GenCheckRegex(`^\d{4}$`), day04GenCheckRange(2010, 2020)},
		"eyr": {day04GenCheckRegex(`^\d{4}$`), day04GenCheckRange(2020, 2030)},
		"hgt": {day04GenCheckRegex(`^\d{2,3}\w{2}$`), day04CheckHgt},
		"hcl": {day04GenCheckRegex(`^#[0-9a-f]{6}$`)},
		"ecl": {day04GenCheckRegex(`^(amb|blu|brn|gry|grn|hzl|oth)$`)},
		"pid": {day04GenCheckRegex(`^\d{9}$`)},
	})
}

func day04CountCheckers(lines []string, checkers map[string][]day04Checker) int {
	count := 0
	for _, passport := range day04ExtractPassports(lines) {
		if day04ValidatePassport(passport, checkers) {
			count++
		}
	}
	return count
}

func day04CheckIsPresent(val string) bool {
	return val != ""
}

func day04CheckHgt(val string) bool {
	if strings.HasSuffix(val, "cm") {
		num := lib.ParseInt(strings.TrimSuffix(val, "cm"))
		return num >= 150 && num <= 193
	}
	if strings.HasSuffix(val, "in") {
		num := lib.ParseInt(strings.TrimSuffix(val, "in"))
		return num >= 59 && num <= 76
	}
	return false
}

func day04GenCheckRange(min, max int) day04Checker {
	return func(val string) bool {
		num := lib.ParseInt(val)
		return num >= min && num <= max
	}
}

func day04GenCheckRegex(pattern string) day04Checker {
	return func(val string) bool {
		re := regexp.MustCompile(pattern)
		return re.MatchString(val)
	}
}

func day04ExtractPassports(lines []string) []map[string]string {
	passports := lib.SplitGroups(lines)

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

func day04ValidatePassport(passport map[string]string, check map[string][]day04Checker) bool {
	for field, checkers := range check {
		for _, checker := range checkers {
			if !checker(passport[field]) {
				return false
			}
		}
	}
	return true
}
