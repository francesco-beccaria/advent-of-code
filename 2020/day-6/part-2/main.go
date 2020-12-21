package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 6
	input      = utils.GetInput(year, day)
	inputArray = strings.Split(input, "\n\n")
)

func main() {
	sum := 0

	for _, s := range inputArray {
		sum += countCommonChars(s)
	}

	fmt.Println(sum)
}

func countCommonChars(s string) int {
	fields := strings.Fields(s)
	fieldsCount := len(fields)
	commonCharsCount := 0

	chars := map[rune]int{}
	for _, item := range fields {
		for _, char := range item {
			chars[char]++
		}
	}

	for _, count := range chars {
		if count == fieldsCount {
			commonCharsCount++
		}
	}
	return commonCharsCount
}
