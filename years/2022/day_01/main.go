package day_01

import (
	"fmt"
	"slices"
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
	sumOfCalories := 0
	maxResult := 0
	// Iterate over the input
	for _, calorie := range input {
		if calorie == "" {
			if maxResult < sumOfCalories {
				maxResult = sumOfCalories
			}
			sumOfCalories = 0
			continue
		}
		calorieInt, err := strconv.Atoi(calorie)
		if err != nil {
			fmt.Println("Error converting calorie:", err)
			continue
		}
		sumOfCalories += calorieInt
	}

	// Add the last elf's calories if input doesn't end with an empty string
	if sumOfCalories > 0 {
		if maxResult < sumOfCalories {
			maxResult = sumOfCalories
		}
	}

	return strconv.Itoa(maxResult)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	elvesWithCalories := []int{}
	sumOfCalories := 0

	// Iterate over the input
	for _, calorie := range input {
		if calorie == "" {
			elvesWithCalories = append(elvesWithCalories, sumOfCalories)
			sumOfCalories = 0
			continue
		}
		calorieInt, err := strconv.Atoi(calorie)
		if err != nil {
			fmt.Println("Error converting calorie:", err)
			continue
		}
		sumOfCalories += calorieInt
	}

	// Add the last elf's calories if input doesn't end with an empty string
	if sumOfCalories > 0 {
		elvesWithCalories = append(elvesWithCalories, sumOfCalories)
	}

	slices.Sort(elvesWithCalories)
	n := len(elvesWithCalories)

	if n >= 3 {
		lastThree := elvesWithCalories[n-3:]
		result := 0

		for i := range lastThree {
			result += lastThree[i]
		}
		return strconv.Itoa(result)
	}
	return ""
}
