package day_02

import (
	"fmt"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	result := 0
	horizontalPosition := 0
	depth := 0

	for _, row := range input {
		parts := strings.Split(row, " ")

		moveCommand := parts[0]
		moveAmount, _ := strconv.Atoi(parts[1])

		switch moveCommand {
		case "forward":
			horizontalPosition += moveAmount
		case "down":
			depth += moveAmount
		case "up":
			depth -= moveAmount
		}
	}

	result = horizontalPosition * depth

	finalResult := strconv.Itoa(result)
	return finalResult
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	result := 0
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, row := range input {
		parts := strings.Split(row, " ")

		moveCommand := parts[0]
		moveAmount, _ := strconv.Atoi(parts[1])

		switch moveCommand {
		case "forward":
			horizontalPosition += moveAmount
			depth += aim * moveAmount
		case "down":
			aim += moveAmount
		case "up":
			aim -= moveAmount
		}
	}

	result = horizontalPosition * depth

	finalResult := strconv.Itoa(result)
	return finalResult
}
