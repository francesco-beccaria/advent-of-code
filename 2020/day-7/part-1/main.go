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
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})

	myBagColour = "shiny gold"
)

func main() {
	bagRules := parseBagRules(inputArray)

	myBagContainers := whoContains(myBagColour, bagRules)

	fmt.Println(len(myBagContainers))
}

func parseBagRules(unparsed []string) map[string]map[string]int {
	parsedBagRules := map[string]map[string]int{}

	for _, s := range unparsed {
		innerOuter := strings.Split(s, " bags contain ")
		outer := innerOuter[0]
		inner := strings.ReplaceAll(strings.ReplaceAll(innerOuter[1], " bags", ""), " bag", "")

		inside := strings.Split(inner, ", ")

		parsedBagRules[outer] = map[string]int{}

		for _, insideBags := range inside {
			if insideBags == "no other." {
				parsedBagRules[outer] = nil
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

func whoContains(colour string, bagRules map[string]map[string]int) map[string]struct{} {
	containers := map[string]struct{}{}

	for outer, inner := range bagRules {
		if _, ok := inner[colour]; ok {
			containers[outer] = struct{}{}
			for c := range whoContains(outer, bagRules) {
				containers[c] = struct{}{}
			}
		}
	}

	return containers
}
