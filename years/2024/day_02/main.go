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

func isRowValid(parts []string) bool {
	incrementor := 0
	for i := 0; i < len(parts)-1; i++ {
		num1, _ := strconv.Atoi(parts[i])
		num2, _ := strconv.Atoi(parts[i+1])

		if i == 0 {
			if num2-num1 < -3 && num2-num1 > 3 {
				return false
			}
			incrementor = num2 - num1
		}

		if num2 == num1 {
			return false
		}

		if incrementor > 0 {
			if num1 > num2 || num2-num1 > 3 {
				return false
			}
		}

		if incrementor < 0 {
			if num2 > num1 || num2-num1 < -3 {
				return false
			}
		}
	}
	return true
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	result := 0
	for _, row := range input {
		parts := strings.Fields(row)
		if isRowValid(parts) {
			result++
		}
	}

	return strconv.Itoa(result)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	result := 0
	for _, row := range input {
		parts := strings.Fields(row)
		if !isRowValid(parts) {
			for i := 0; i < len(parts); i++ {
				tempParts := make([]string, 0)
				tempParts = append(tempParts, parts[:i]...)
				tempParts = append(tempParts, parts[i+1:]...)

				if isRowValid(tempParts) {
					result++
					break
				}
			}

		} else {
			result++
		}
	}

	return strconv.Itoa(result)
}
