package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 4
	input      = utils.GetInput(year, day)
	inputArray = strings.Split(input, "\n\n")

	requiredFields = map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
		"cid": false,
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
		vals := strings.Split(item, ":")
		passport[vals[0]] = vals[1]
	}

	return passport
}

func isValidPassport(passport map[string]string) bool {
	for field, required := range requiredFields {
		if _, ok := passport[field]; !ok && required {
			return false
		}
	}
	return true
}
