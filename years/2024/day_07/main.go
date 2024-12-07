package day_07

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
	var count float64 = 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		target, _ := strconv.ParseFloat(parts[0], 64)

		numbersStr := strings.Fields(parts[1])
		numbers := make([]float64, len(numbersStr))
		for i, num := range numbersStr {
			numbers[i], _ = strconv.ParseFloat(num, 64)
		}

		resultFound := FindCombinationWithTarget(numbers, target, false)
		if resultFound != nil {
			count += target
		}
	}

	return strconv.FormatFloat(count, 'f', 0, 64)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	var count float64 = 0

	for _, line := range input {
		parts := strings.Split(line, ": ")
		target, _ := strconv.ParseFloat(parts[0], 64)

		numbersStr := strings.Fields(parts[1])
		numbers := make([]float64, len(numbersStr))
		for i, num := range numbersStr {
			numbers[i], _ = strconv.ParseFloat(num, 64)
		}

		resultFound := FindCombinationWithTarget(numbers, target, true)
		if resultFound != nil {
			count += target
		}
	}

	return strconv.FormatFloat(count, 'f', 0, 64)
}

type Result struct {
	expression string
	result     float64
}

func FindCombinationWithTarget(numbers []float64, targetValue float64, allowConcatenation bool) *Result {
	if len(numbers) == 0 {
		return nil
	}
	return GenerateCombinationsRecursive(numbers, 1, numbers[0], targetValue, allowConcatenation)
}

func GenerateCombinationsRecursive(numbers []float64, index int, currentValue float64, targetValue float64, allowConcatenation bool) *Result {
	if index == len(numbers) {
		if currentValue == targetValue {
			return &Result{
				result: currentValue,
			}
		}
		return nil
	}

	nextNumber := numbers[index]

	if result := GenerateCombinationsRecursive(
		numbers,
		index+1,
		currentValue+nextNumber,
		targetValue,
		allowConcatenation,
	); result != nil {
		return result
	}

	if result := GenerateCombinationsRecursive(
		numbers,
		index+1,
		currentValue*nextNumber,
		targetValue,
		allowConcatenation,
	); result != nil {
		return result
	}

	if allowConcatenation {
		concatenated, _ := strconv.ParseFloat(
			fmt.Sprintf("%.0f%.0f", currentValue, nextNumber),
			64,
		)
		result := GenerateCombinationsRecursive(
			numbers,
			index+1,
			concatenated,
			targetValue,
			allowConcatenation,
		)

		return result
	}
	return nil
}
