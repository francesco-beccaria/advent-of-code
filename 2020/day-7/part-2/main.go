package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"advent-of-code/utils"
)

var (
	year       = 2020
	day        = 7
	input      = utils.GetInput(year, day)
	inputArray = strings.Split(input, "\n")

	myBagColour = "shiny gold"
)

func main() {
	bagRules := parseBagRules(inputArray)

	myBagContains := contains(myBagColour, bagRules)

	fmt.Println(myBagContains)
}

func parseBagRules(unparsed []string) map[string]map[string]int {
	parsedBagRules := map[string]map[string]int{}

	for _, s := range unparsed {
		if s == "" {
			break
		}

		innerOuter := strings.Split(s, " bags contain ")
		outer := innerOuter[0]
		inner := strings.ReplaceAll(strings.ReplaceAll(innerOuter[1], " bags", ""), " bag", "")

		inside := strings.Split(inner, ", ")

		parsedBagRules[outer] = map[string]int{}

		for _, insideBags := range inside {
			if insideBags == "no other." {
				parsedBagRules[outer] = map[string]int{}
				continue
			}
			insideNum, err := strconv.Atoi(strings.TrimRightFunc(insideBags, func(c rune) bool {
				return !unicode.IsNumber(c)
			}))
			if err != nil {
				log.Fatalf("failed converting to int: %s", insideBags)
			}
			insideColour := strings.TrimFunc(insideBags, func(c rune) bool {
				return !unicode.IsLetter(c)
			})
			parsedBagRules[outer][insideColour] = insideNum
		}
	}

	return parsedBagRules
}

func contains(colour string, bagRules map[string]map[string]int) int {
	bags := 0

	if len(bagRules[colour]) == 0 {
		return bags
	}

	for c, i := range bagRules[colour] {
		bags += i + i*contains(c, bagRules)
	}

	return bags
}
