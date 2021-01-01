package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 12
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})

	directions = []string{
		"North",
		"East",
		"South",
		"West",
	}
	currentDirection = "East"
)

func main() {
	fmt.Println(finalPosition(inputArray))
}

func finalPosition(input []string) int {
	currentPositionX := 0
	currentPositionY := 0

	for _, v := range input {
		letter := v[0]
		val, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatalf("failed converting to int: %s\n", v[1:])
		}
		if letter == 'F' {
			letter = currentDirection[0]
		}
		switch letter {
		case 'N':
			currentPositionY += val
		case 'S':
			currentPositionY -= val
		case 'E':
			currentPositionX += val
		case 'W':
			currentPositionX -= val
		case 'R':
			currentDirection = turnRight(currentDirection, val)
		case 'L':
			currentDirection = turnLeft(currentDirection, val)
		}
	}

	return int(math.Abs(float64(currentPositionX)) + math.Abs(float64(currentPositionY)))
}

func turnRight(originalDirection string, degree int) string {
	var newDirection string

	for i, dir := range directions {
		if dir == originalDirection {
			newDirection = directions[(i+degree/90)%4]
			break
		}
	}

	return newDirection
}

func turnLeft(originalDirection string, degree int) string {
	return turnRight(originalDirection, (360 - degree))
}
