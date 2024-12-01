package day_01

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
	leftList := []int{}
	rightList := []int{}

	for _, row := range input {
		parts := strings.Fields(row)

		num1, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		leftList = insertSorted(leftList, num1)
		rightList = insertSorted(rightList, num2)
	}

	for i := 0; i < len(leftList); i++ {
		result += absInt(leftList[i] - rightList[i])
	}
	return strconv.Itoa(result)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {

	result := 0
	leftList := []int{}
	occurrencesInRightList := make(map[int]int)

	for _, row := range input {
		parts := strings.Fields(row)

		num1, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		leftList = append(leftList, num1)
		occurrencesInRightList[num2]++

	}

	for _, number := range leftList {
		result += number * occurrencesInRightList[number]
	}

	return strconv.Itoa(result)
}

func insertSorted(list []int, num int) []int {
	inserted := false
	for i := 0; i < len(list); i++ {
		if num < list[i] {
			// Insert at the correct position
			list = append(list[:i], append([]int{num}, list[i:]...)...)
			inserted = true
			break
		}
	}
	if !inserted {
		list = append(list, num)
	}
	return list
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
