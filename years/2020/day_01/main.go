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

	hashMaps := map[int]int{}
	goalToReach := 2020

	for i := range input {
		newValueToBeAdded, _ := strconv.Atoi(input[i])
		value, exists := hashMaps[goalToReach-newValueToBeAdded]

		if exists {
			return strconv.Itoa(value * (goalToReach - newValueToBeAdded))
		} else {
			hashMaps[newValueToBeAdded] = goalToReach - newValueToBeAdded
		}
	}
	return ""
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {

	hashMaps := make(map[int]Pair)
	goalToReach := 2020

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if i != j {
				newValueToBeAdded, _ := strconv.Atoi(input[i])
				newValueToBeAdded2, _ := strconv.Atoi(input[j])

				valueToBeAddedOnMap := goalToReach - newValueToBeAdded - newValueToBeAdded2

				if valueToBeAddedOnMap > 0 {
					hashMaps[valueToBeAddedOnMap] = Pair{First: newValueToBeAdded, Second: newValueToBeAdded2}
				}

			}
		}
	}

	for i := range input {
		valueToBeChecked, _ := strconv.Atoi(input[i])

		value, exists := hashMaps[valueToBeChecked]

		if exists {
			return strconv.Itoa(valueToBeChecked * value.First * value.Second)
		}
	}
	return ""
}

type Pair struct {
	First  int
	Second int
}
