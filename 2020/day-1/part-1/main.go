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
		if _, ok := intMap[(2020 - i)]; ok {
			fmt.Println(i * (2020 - i))
			os.Exit(0)
		}
	}
}
