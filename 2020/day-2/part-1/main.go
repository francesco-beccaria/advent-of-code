package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 2
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, splitFunc)

	splitFunc = func(c rune) bool {
		return c == '\n'
	}
)

func main() {
	validPasswords := 0

	for _, s := range inputArray {
		minMax := strings.Split(s, " ")[0]
		min, _ := strconv.Atoi(strings.Split(minMax, "-")[0])
		max, _ := strconv.Atoi(strings.Split(minMax, "-")[1])
		letter := strings.Split(strings.Split(s, ":")[0], " ")[1]
		password := strings.Split(s, ": ")[1]

		occurrencesOfLetter := len(strings.Split(password, letter)) - 1
		if occurrencesOfLetter >= min && occurrencesOfLetter <= max {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
}
