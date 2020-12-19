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
	mySeatID := 0
	seatIDs := map[int]struct{}{}

	for _, seatSpec := range inputArray {
		seatIDs[getSeatID(seatSpec)] = struct{}{}
	}

	for seatID := range seatIDs {
		_, okabove1 := seatIDs[seatID+1]
		_, okabove2 := seatIDs[seatID+2]
		isAbove := okabove2 && !okabove1

		_, okbelow1 := seatIDs[seatID-1]
		_, okbelow2 := seatIDs[seatID-2]
		isBelow := okbelow2 && !okbelow1

		if isAbove {
			mySeatID = seatID + 1
			break
		}
		if isBelow {
			mySeatID = seatID - 1
			break
		}
	}
	fmt.Println(mySeatID)
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
