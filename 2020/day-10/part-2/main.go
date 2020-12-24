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

	fmt.Println(countPaths(inputArrayInt))
}

func countPaths(input []int) int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	input = append(input, input[len(input)-1]+3)

	cache := map[int]int{}
	cache[0] = 1

	for _, v := range input {
		cache[v] = cache[v-3] + cache[v-2] + cache[v-1]
	}

	return cache[input[len(input)-1]]
}
