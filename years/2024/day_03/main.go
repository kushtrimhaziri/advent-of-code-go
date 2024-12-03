package day_03

import (
	"fmt"
	"regexp"
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
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0

	for _, row := range input {
		matches1 := re.FindAllStringSubmatch(row, -1)
		fmt.Println("Matches found:", len(matches1))
		for _, match := range matches1 {
			if len(match) == 3 {
				num1, err1 := strconv.Atoi(match[1])
				num2, err2 := strconv.Atoi(match[2])
				if err1 == nil && err2 == nil {
					result += num1 * num2
				}
			}
		}
	}
	return strconv.Itoa(result)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	result := 0
	isEnabled := true

	for _, row := range input {
		matches := re.FindAllStringSubmatch(row, -1)

		fmt.Println("Matches found:")
		for _, match := range matches {
			if match[1] != "" && match[2] != "" {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				if isEnabled {
					result += num1 * num2
				}
			} else if match[0] == "do()" {
				isEnabled = true
			} else if match[0] == "don't()" {
				isEnabled = false
			}
		}
	}
	return strconv.Itoa(result)
}
