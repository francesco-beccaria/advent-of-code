package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 1
	input      = utils.GetInput(year, day)
	inputArray = strings.Fields(input)
)

func main() {
	intMap := make(map[int]bool)
	for _, s := range inputArray {
		i, _ := strconv.Atoi(s)
		intMap[i] = true
	}

	for i := range intMap {
		for j := range intMap {
			if _, ok := intMap[(2020 - i - j)]; ok && j != i {
				fmt.Println(i * j * (2020 - i - j))
				os.Exit(0)
			}
		}
	}
}
