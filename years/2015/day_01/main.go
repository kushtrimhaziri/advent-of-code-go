package day_01

import (
	"fmt"
	"strconv"
)

var Parentheses map[string]int = map[string]int{
	"(": 0,
	")": 0,
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {

		fmt.Printf("Part one: %v\n", Part1(input[0]))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input[0]))
	}
}

// Part1 solves the first part of the exercise
func Part1(input string) string {

	for _, elem := range input {
		Parentheses[string(elem)]++
	}
	positionOfChar := Parentheses["("] - Parentheses[")"]

	return strconv.Itoa(positionOfChar)
}

// Part2 solves the second part of the exercise
func Part2(input string) string {
	positionOfChar := 1

	for i, elem := range input {
		if string(elem) == "(" {
			positionOfChar++
		} else {
			positionOfChar--
		}
		if positionOfChar < 0 {
			return strconv.Itoa(i)
		}
	}
	return ""
}
