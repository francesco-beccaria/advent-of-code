package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

var (
	year     = 2020
	day      = 4
	input    = utils.GetInput(year, day)
	strArray = strings.Split(input, "\n\n")
)

func main() {
	validPassports := 0

	for _, s := range strArray {
		if isValidPassport(newPassport(s)) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func newPassport(s string) map[string]string {
	fields := strings.Fields(s)

	passport := make(map[string]string)

	for _, item := range fields {
		vals := strings.Split(item, ":")
		passport[vals[0]] = vals[1]
	}

	return passport
}

func isValidPassport(passport map[string]string) bool {
	if _, ok := passport["byr"]; ok {
		if _, ok := passport["iyr"]; ok {
			if _, ok := passport["eyr"]; ok {
				if _, ok := passport["hgt"]; ok {
					if _, ok := passport["hcl"]; ok {
						if _, ok := passport["ecl"]; ok {
							if _, ok := passport["pid"]; ok {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}
