package day_01

import (
	"fmt"
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
	var finalSum int = 0

	for i := range input {

		valueToBeAdded, err := strconv.Atoi(input[i])
		if err != nil {
			return ""
		}

		finalSum += valueToBeAdded
	}

	return strconv.Itoa(finalSum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	//setOfFrequencies := make(map[int]struct{})
	//currentSum := 0
	//indexOfStart := 0
	//expectedFrequency, err := strconv.Atoi(input[0])
	//if err != nil {
	//	return ""
	//}
	//
	//for _, exists := setOfFrequencies[expectedFrequency]; exists != true
	//{
	//	currentSum+=
	//}
	//for i := range input {
	//	return ""

	return ""
}
