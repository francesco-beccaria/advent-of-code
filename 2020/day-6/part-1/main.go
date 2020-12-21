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
		sum += countDistinctChars(s)
	}

	fmt.Println(sum)
}

func countDistinctChars(s string) int {
	fields := strings.Fields(s)

	distinctChars := map[rune]struct{}{}
	for _, item := range fields {
		for _, char := range item {
			distinctChars[char] = struct{}{}
		}
	}

	return len(distinctChars)
}
