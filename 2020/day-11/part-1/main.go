package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 11
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})

	columns = len(inputArray[0])
	rows    = len(inputArray)

	neighbours = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
)

func main() {
	occupiedSeats := countOccupiedSeats(inputArray)

	fmt.Println(occupiedSeats)
}

func countOccupiedSeats(input []string) int {
	changed := false
	output := make([]string, rows)
	occupied := 0

	for y, line := range input {
		outputLine := make([]rune, columns)
		for x, char := range line {
			occupiedNeighbours := 0
			for _, n := range neighbours {
				nX := x + n[0]
				nY := y + n[1]
				if 0 <= nX && nX < columns && 0 <= nY && nY < rows && input[nY][nX] == '#' {
					occupiedNeighbours++
				}
			}
			switch char {
			case 'L':
				if occupiedNeighbours == 0 {
					outputLine[x] = rune('#')
					changed = true
				} else {
					outputLine[x] = rune('L')
				}
			case '#':
				if occupiedNeighbours >= 4 {
					outputLine[x] = rune('L')
					changed = true
				} else {
					outputLine[x] = rune('#')
					occupied++
				}
			case '.':
				outputLine[x] = rune('.')
			}
		}
		output[y] = string(outputLine)
	}
	if !changed {
		return occupied
	}
	return countOccupiedSeats(output)
}
