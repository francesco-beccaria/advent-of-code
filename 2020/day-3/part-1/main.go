package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 3
	input      = utils.GetInput(year, day)
	inputArray = strings.Fields(input)
)

func main() {
	fmt.Println(countTrees(inputArray, 3, 1))
}

func countTrees(a []string, x int, y int) int {
	lineLength := len(a[0])
	trees := 0
	posHor := 0
	posVer := 0

	for _, s := range a {
		if posVer%y == 0 {
			if string(s[posHor]) == "#" {
				trees++
			}
			posHor = (posHor + x) % lineLength
		}
		posVer++
	}

	return trees
}
