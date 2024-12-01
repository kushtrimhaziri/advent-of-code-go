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
		parts := strings.Split(row, "x")

		length, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		width, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		height, _ := strconv.Atoi(strings.TrimSpace(parts[2]))

		result += 2*(length*width+length*height+width*height) + min(length*width, length*height, width*height)

	}
	finalResult := strconv.Itoa(result)
	return finalResult
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	finalSum := 0

	for _, row := range input {
		parts := strings.Split(row, "x")

		length, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		width, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		height, _ := strconv.Atoi(strings.TrimSpace(parts[2]))

		smallestPerimeter := min(
			2*(length+width),
			2*(width+height),
			2*(length+height),
		)

		finalSum += smallestPerimeter + (length * width * height)

	}
	finalResult := strconv.Itoa(finalSum)
	return finalResult
}
