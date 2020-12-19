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
		firstSecond := strings.Split(s, " ")[0]
		first, _ := strconv.Atoi(strings.Split(firstSecond, "-")[0])
		second, _ := strconv.Atoi(strings.Split(firstSecond, "-")[1])
		letter := strings.Split(strings.Split(s, ":")[0], " ")[1]
		password := strings.Split(s, ": ")[1]

		if (string(password[first-1]) == letter) != (string(password[second-1]) == letter) {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
}
