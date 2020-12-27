package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 9
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
	value, err := findFirstWrongValue(inputArrayInt)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	min, max := findContiguous(inputArrayInt, value, 0, 0)

	fmt.Println(min + max)
}

func findFirstWrongValue(values []int) (int, error) {
	for i := 25; i < len(values); i++ {
		if !isValueOK(values[i], values[i-25:i]) {
			return values[i], nil
		}
	}
	return 0, fmt.Errorf("failed to find a wrong value")
}

func isValueOK(value int, previousValues []int) bool {
	previousValuesMap := map[int]int{}
	for _, v := range previousValues {
		previousValuesMap[v] = value - v
	}
	for a, b := range previousValuesMap {
		if _, ok := previousValuesMap[b]; ok && a != b {
			return true
		}
	}
	return false
}

func findContiguous(input []int, target int, start int, end int) (int, int) {
	sum := 0
	min := input[start]
	max := min
	for _, v := range input[start:end] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
		if sum == target {
			return min, max
		}
		if sum > target {
			return findContiguous(input, target, start+1, end)
		}
	}
	return findContiguous(input, target, start, end+1)
}
