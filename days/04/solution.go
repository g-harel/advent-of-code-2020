package solution

import (
	"regexp"
	"strings"

	"lib"
)

type Checker func(string) bool

func Part1() int {
	return countCheckers(lib.ReadLines("input.txt"), map[string][]Checker{
		"byr": {checkIsPresent},
		"iyr": {checkIsPresent},
		"eyr": {checkIsPresent},
		"hgt": {checkIsPresent},
		"hcl": {checkIsPresent},
		"ecl": {checkIsPresent},
		"pid": {checkIsPresent},
	})
}

func Part2() int {
	return countCheckers(lib.ReadLines("input.txt"), map[string][]Checker{
		"byr": {genCheckRegex(`^\d{4}$`), genCheckRange(1920, 2002)},
		"iyr": {genCheckRegex(`^\d{4}$`), genCheckRange(2010, 2020)},
		"eyr": {genCheckRegex(`^\d{4}$`), genCheckRange(2020, 2030)},
		"hgt": {genCheckRegex(`^\d{2,3}\w{2}$`), checkHgt},
		"hcl": {genCheckRegex(`^#[0-9a-f]{6}$`)},
		"ecl": {genCheckRegex(`^(amb|blu|brn|gry|grn|hzl|oth)$`)},
		"pid": {genCheckRegex(`^\d{9}$`)},
	})
}

func countCheckers(lines []string, checkers map[string][]Checker) int {
	count := 0
	for _, passport := range extractPassports(lines) {
		if ValidatePassport(passport, checkers) {
			count++
		}
	}
	return count
}

func checkIsPresent(val string) bool {
	return val != ""
}

func checkHgt(val string) bool {
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

func genCheckRange(min, max int) Checker {
	return func(val string) bool {
		num := lib.ParseInt(val)
		return num >= min && num <= max
	}
}

func genCheckRegex(pattern string) Checker {
	return func(val string) bool {
		re := regexp.MustCompile(pattern)
		return re.MatchString(val)
	}
}

func extractPassports(lines []string) []map[string]string {
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

func ValidatePassport(passport map[string]string, check map[string][]Checker) bool {
	for field, checkers := range check {
		for _, checker := range checkers {
			if !checker(passport[field]) {
				return false
			}
		}
	}
	return true
}
