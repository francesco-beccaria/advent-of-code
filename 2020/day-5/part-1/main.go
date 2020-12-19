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
	day        = 5
	input      = utils.GetInput(year, day)
	inputArray = strings.Fields(input)
)

func main() {
	maxSeatID := 0

	for _, seatSpec := range inputArray {
		seatID := getSeatID(seatSpec)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println(maxSeatID)
}

func getSeatID(spec string) int {
	rowPart := spec[:7]
	columnPart := spec[7:]

	rowBin := strings.ReplaceAll(strings.ReplaceAll(rowPart, "F", "0"), "B", "1")
	columnBin := strings.ReplaceAll(strings.ReplaceAll(columnPart, "L", "0"), "R", "1")

	row, err := strconv.ParseInt(rowBin, 2, 64)
	if err != nil {
		log.Fatalf("failed converting binary string to decimal: %s", err)
	}

	column, err := strconv.ParseInt(columnBin, 2, 64)
	if err != nil {
		log.Fatalf("failed converting binary string to decimal: %s", err)
	}

	return int(row*8 + column)
}
