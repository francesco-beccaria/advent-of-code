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
	day        = 13
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})
)

func main() {
	earliestDeparture, err := strconv.Atoi(inputArray[0])
	if err != nil {
		log.Fatalf("failed converting to int: %s", inputArray[0])
	}

	busIDs := strings.Split(inputArray[1], ",")
	buses := make(map[int]int, len(busIDs))
	for _, b := range busIDs {
		if b == "x" {
			continue
		}
		bInt, err := strconv.Atoi(b)
		if err != nil {
			log.Fatalf("failed converting to int: %s", b)
		}
		busPreviousArrival := earliestDeparture % bInt
		buses[bInt] = bInt - busPreviousArrival
	}

	earliestBusID := 0
	earliestBusWait := 0
	for b, w := range buses {
		if earliestBusID == 0 || w < earliestBusWait {
			earliestBusID = b
			earliestBusWait = w
		}
	}

	fmt.Println(earliestBusID * earliestBusWait)
}
