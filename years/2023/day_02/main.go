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
	for _, value := range input {
		parts := strings.SplitN(value, ":", 2)
		gameNumber := strings.SplitN(parts[0], " ", 2)[1]
		rounds := strings.Split(parts[1], ";")

		isValid := true
		for _, value := range rounds {
			var Numbers map[string]int = map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}
			sets := strings.Split(value, ",")
			for _, value := range sets {
				numberAndColor := strings.Fields(value)
				number, err := strconv.Atoi(numberAndColor[0])
				if err == nil {
					Numbers[numberAndColor[1]] = Numbers[numberAndColor[1]] - number
				}

				if Numbers[numberAndColor[1]] < 0 {
					isValid = false
					break
				}
			}
		}

		if isValid {
			gameNumberAsInt, err := strconv.Atoi(gameNumber)
			if err == nil {
				result += gameNumberAsInt
			}
		}
	}
	finalResult := strconv.Itoa(result)
	return finalResult
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	finalSum := 0
	for _, value := range input {
		var Numbers map[string]int = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		parts := strings.SplitN(value, ":", 2)
		rounds := strings.Split(parts[1], ";")

		for _, value := range rounds {
			sets := strings.Split(value, ",")
			for _, value := range sets {
				numberAndColor := strings.Fields(value)
				number, err := strconv.Atoi(numberAndColor[0])
				if err == nil {
					if number > Numbers[numberAndColor[1]] {
						Numbers[numberAndColor[1]] = number
					}
				}
			}
		}

		result := 1
		for _, value := range Numbers {
			result *= value
		}
		finalSum += result
	}

	finalResult := strconv.Itoa(finalSum)
	return finalResult
}
