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

	for _, row := range input {
		parts := strings.Fields(row)

		intRow := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			intRow = append(intRow, num)

		}

		result += max(intRow...) - min(intRow...)

	}

	finalResult := strconv.Itoa(result)
	return finalResult
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	finalSum := 0

	for _, row := range input {
		parts := strings.Fields(row)

		intRow := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			intRow = append(intRow, num)

		}

		pairs := findDivisiblePairs(intRow)
		for _, pair := range pairs {
			finalSum += pair[0] / pair[1]
		}

	}

	finalResult := strconv.Itoa(finalSum)
	return finalResult
}

func max(nums ...int) int {
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func min(nums ...int) int {
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func findDivisiblePairs(nums []int) [][2]int {
	var pairs [][2]int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[j] != 0 && nums[i]%nums[j] == 0 {
				pairs = append(pairs, [2]int{nums[i], nums[j]})
			}
		}
	}

	return pairs
}
