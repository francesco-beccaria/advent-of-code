package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

type instruction struct {
	command string
	value   int
	pos     int
}

var (
	year       = 2020
	day        = 10
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})
)

func main() {
	inputArrayInt := make([]int, len(inputArray))
	for i, v := range inputArray {
		var err error
		if inputArrayInt[i], err = strconv.Atoi(v); err != nil {
			log.Fatalf("failed converting to int: %s", v)
		}
	}
	oneJoltDiff, threeJoltDiff := countOneThreeJoltDiffs(inputArrayInt)

	fmt.Println(oneJoltDiff * threeJoltDiff)
}

func countOneThreeJoltDiffs(input []int) (int, int) {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	oneJoltDiff := 0
	threeJoltDiff := 1
	previousJoltRate := 0

	for _, v := range input {
		if v-previousJoltRate == 1 {
			oneJoltDiff++
		}
		if v-previousJoltRate == 3 {
			threeJoltDiff++
		}
		previousJoltRate = v
	}

	return oneJoltDiff, threeJoltDiff
}
