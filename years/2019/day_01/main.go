package day_01

import (
	"fmt"
	"math"
	"strconv"
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

	var finalResult = 0

	for i := range input {
		newValueToBeAdded, _ := strconv.Atoi(input[i])

		finalResult += int(math.Floor(float64(newValueToBeAdded)/3)) - 2
	}
	return strconv.Itoa(finalResult)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	var finalResult = 0

	for i := range input {
		newValueToBeAdded, _ := strconv.Atoi(input[i])
		newValueToBeRemoved, _ := strconv.Atoi(input[i])
		totalSumToBeAdded := 0
		for newValueToBeAdded > 0 {
			totalSumToBeAdded += newValueToBeAdded
			newValueToBeAdded = int(math.Floor(float64(newValueToBeAdded)/3)) - 2
		}
		finalResult += totalSumToBeAdded - newValueToBeRemoved
	}

	return strconv.Itoa(finalResult)
}
