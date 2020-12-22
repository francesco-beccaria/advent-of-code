package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

type instruction struct {
	command string
	value   int
	pos     int
}

var (
	year       = 2020
	day        = 8
	input      = utils.GetInput(year, day)
	inputArray = strings.FieldsFunc(input, func(c rune) bool {
		return c == '\n'
	})
)

func main() {
	instructions := parseInstructions(inputArray)

	acc, _, err := executeInstructions(instructions)
	if err != nil {
		log.Printf("info: %s", err)
	}

	fmt.Println(acc)
}

func parseInstructions(unparsed []string) []instruction {
	parsed := make([]instruction, len(unparsed))

	for i, s := range unparsed {
		splitS := strings.Fields(s)
		command := splitS[0]
		value, err := strconv.Atoi(splitS[1])
		if err != nil {
			log.Fatalf("failed converting to int: %s", splitS[1])
		}
		parsed[i] = instruction{
			command: command,
			value:   value,
			pos:     i,
		}
	}

	return parsed
}

func executeInstructions(instructions []instruction) (int, []instruction, error) {
	executedInstructions := []instruction{}

	previousInstruction := instruction{}
	acc := 0
	currentPos := 0

	for currentPos < len(instructions) {
		for _, v := range executedInstructions {
			if currentPos == v.pos {
				return acc, executedInstructions, fmt.Errorf("found loop at %v", previousInstruction)
			}
		}

		currentInstruction := instructions[currentPos]
		previousInstruction = currentInstruction

		executedInstructions = append(executedInstructions, currentInstruction)

		switch currentInstruction.command {
		case "acc":
			acc += currentInstruction.value
			currentPos++
		case "jmp":
			currentPos += currentInstruction.value
		case "nop":
			currentPos++
		}
	}

	return acc, nil, nil
}
