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

	waypoint = map[string]int{"x": 10, "y": 1}
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
		switch letter {
		case 'N':
			waypoint["y"] += val
		case 'S':
			waypoint["y"] -= val
		case 'E':
			waypoint["x"] += val
		case 'W':
			waypoint["x"] -= val
		case 'R':
			waypoint = rotateWaypointClockwise(waypoint, val)
		case 'L':
			waypoint = rotateWaypointCounterClockwise(waypoint, val)
		case 'F':
			currentPositionX = currentPositionX + (val * waypoint["x"])
			currentPositionY = currentPositionY + (val * waypoint["y"])
		}
	}

	return int(math.Abs(float64(currentPositionX)) + math.Abs(float64(currentPositionY)))
}

func rotateWaypointClockwise(currentWaypoint map[string]int, degree int) map[string]int {
	newWaypoint := map[string]int{}

	for i := 0; i < degree/90; i++ {
		newWaypoint["x"] = currentWaypoint["y"]
		newWaypoint["y"] = -currentWaypoint["x"]
		currentWaypoint["x"] = newWaypoint["x"]
		currentWaypoint["y"] = newWaypoint["y"]
	}

	return newWaypoint
}

func rotateWaypointCounterClockwise(currentWaypoint map[string]int, degree int) map[string]int {
	return rotateWaypointClockwise(currentWaypoint, (360 - degree))
}
