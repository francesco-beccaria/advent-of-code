package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

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
	if val, ok := passport["byr"]; ok {
		if valInt, _ := strconv.Atoi(val); valInt >= 1920 && valInt <= 2002 {
			if val, ok := passport["iyr"]; ok {
				if valInt, _ := strconv.Atoi(val); valInt >= 2010 && valInt <= 2020 {
					if val, ok := passport["eyr"]; ok {
						if valInt, _ := strconv.Atoi(val); valInt >= 2020 && valInt <= 2030 {
							if val, ok := passport["hgt"]; ok {
								if valInt, _ := strconv.Atoi(val[0 : len(val)-2]); (val[len(val)-2:] == "cm" && valInt >= 150 && valInt <= 193) || (val[len(val)-2:] == "in" && valInt >= 59 && valInt <= 76) {
									if val, ok := passport["hcl"]; ok {
										if b, _ := regexp.MatchString("^#[0-9abcdef]{6}$", val); b {
											if val, ok := passport["ecl"]; ok {
												if val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth" {
													if val, ok := passport["pid"]; ok {
														if b, _ := regexp.MatchString("^[0-9]{9}$", val); b {
															return true
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func main() {
	year := 2020
	day := 4

	input := utils.GetInput(year, day)

	strArray := strings.Split(input, "\n\n")

	validPassports := 0

	for _, s := range strArray {
		if isValidPassport(newPassport(s)) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}
