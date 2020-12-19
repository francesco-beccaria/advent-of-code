package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

type validation func(string) bool

var (
	year       = 2020
	day        = 4
	input      = utils.GetInput(year, day)
	inputArray = strings.Split(input, "\n\n")

	requiredFields = map[string]validation{
		"byr": func(val string) bool {
			year, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return 1920 <= year && year <= 2002
		},
		"iyr": func(val string) bool {
			year, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return 2010 <= year && year <= 2020
		},
		"eyr": func(val string) bool {
			year, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			return 2020 <= year && year <= 2030
		},
		"hgt": func(val string) bool {
			unit := val[len(val)-2:]
			num, err := strconv.Atoi(val[:len(val)-2])
			if err != nil {
				return false
			}
			switch unit {
			case "cm":
				return 150 <= num && num <= 193
			case "in":
				return 59 <= num && num <= 76
			}
			return false
		},
		"hcl": func(val string) bool {
			match, err := regexp.MatchString(`^#[\da-f]{6}$`, val)
			if err != nil {
				return false
			}
			return match
		},
		"ecl": func(val string) bool {
			switch val {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				return true
			}
			return false
		},
		"pid": func(val string) bool {
			match, err := regexp.MatchString(`^\d{9}$`, val)
			if err != nil {
				return false
			}
			return match
		},
	}
)

func main() {
	validPassports := 0

	for _, s := range inputArray {
		if isValidPassport(newPassport(s)) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func newPassport(s string) map[string]string {
	fields := strings.Fields(s)

	passport := map[string]string{}

	for _, item := range fields {
		splitItems := strings.Split(item, ":")
		passport[splitItems[0]] = splitItems[1]
	}

	return passport
}

func isValidPassport(passport map[string]string) bool {
	for field, isValid := range requiredFields {
		value, ok := passport[field]
		if !ok {
			return false
		}
		if !isValid(value) {
			return false
		}
	}
	return true
}
