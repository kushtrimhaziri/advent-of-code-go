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
	firstNumber, _ := strconv.Atoi(input[0])
	nrOfIncrements := 0

	for value := range input {
		numberToBeCompared, _ := strconv.Atoi(input[value])

		if firstNumber < numberToBeCompared {
			nrOfIncrements++
		}

		firstNumber = numberToBeCompared
	}
	return strconv.Itoa(nrOfIncrements)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	firstNumber, _ := strconv.Atoi(input[0])
	secondNumber, _ := strconv.Atoi(input[1])
	thirdNumber, _ := strconv.Atoi(input[2])
	sumOfFirstThreeNumbers := firstNumber + secondNumber + thirdNumber
	nrOfIncrements := 0

	for i := 1; i < len(input)-2; i++ {
		firstNumberToBeCompared, _ := strconv.Atoi(input[i])
		secondNumberToBeCompared, _ := strconv.Atoi(input[i+1])
		thirdNumberToBeCompared, _ := strconv.Atoi(input[i+2])

		sumOfFirstThreeNumbersToBeCompared := firstNumberToBeCompared + secondNumberToBeCompared + thirdNumberToBeCompared

		if sumOfFirstThreeNumbers < sumOfFirstThreeNumbersToBeCompared {
			nrOfIncrements++
		}

		sumOfFirstThreeNumbers = sumOfFirstThreeNumbersToBeCompared
	}
	return strconv.Itoa(nrOfIncrements)
}
