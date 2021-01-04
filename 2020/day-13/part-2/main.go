package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

type bus struct {
	id      *big.Int
	offset  *big.Int
	matched bool
}

var (
	year       = 2020
	day        = 13
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})
)

func main() {
	buses := parseBusses(strings.Split(inputArray[1], ","))

	timestamp := calculateMatchingTimestamp(buses)

	fmt.Println(timestamp.String())
}

func parseBusses(b []string) []bus {
	busses := []bus{}

	for i, v := range b {
		id, err := strconv.Atoi(v)
		if err == nil {
			busses = append(busses, bus{
				id:      big.NewInt(int64(id)),
				offset:  big.NewInt(int64(i)),
				matched: false,
			})
		}
	}

	return busses
}

func calculateMatchingTimestamp(buses []bus) *big.Int {
	mod := big.NewInt(0)
	zero := big.NewInt(0)
	timestamp := big.NewInt(0)
	lcm := buses[0].id
	gcd := big.NewInt(0)
	buses[0].matched = true
	matched := 1

	for matched < len(buses) {
		timestamp.Add(timestamp, lcm)

		for i, bus := range buses {
			if !bus.matched {
				mod.Mod(mod.Add(timestamp, bus.offset), bus.id)
				if mod.Cmp(zero) == 0 {
					buses[i].matched = true
					matched++
					if matched == len(buses) {
						break
					}
					gcd.GCD(nil, nil, lcm, bus.id)
					lcm.Mul(lcm, bus.id)
					lcm.Div(lcm, gcd)
				}
			}
		}
	}

	return timestamp
}
