package day_01

import (
	"fmt"
	"strconv"
)

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
	currentSum := 0

	for i := 0; i < len(input); i++ {
		first := input[i]
		second := input[(i+1)%len(input)]

		if first == second {
			convertedNumberToBeAdded, err := strconv.Atoi(string(first))
			if err != nil {
				return ""
			}
			currentSum += convertedNumberToBeAdded
		}
	}
	return strconv.Itoa(currentSum)
}

// Part2 solves the second part of the exercise
func Part2(input string) string {
	jumper := len(input) / 2
	currentSum := 0

	for i := 0; i < len(input); i++ {
		first := input[i]
		second := input[(i+jumper)%len(input)]

		if first == second {
			convertedNumberToBeAdded, err := strconv.Atoi(string(first))
			if err != nil {
				return ""
			}
			currentSum += convertedNumberToBeAdded
		}
	}
	return strconv.Itoa(currentSum)

}
