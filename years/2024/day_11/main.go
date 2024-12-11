package day_11

import (
	"fmt"
	"math/big"
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
	numberMap := make(map[float64]big.Int)

	numbers := strings.Fields(input[0])
	for _, numStr := range numbers {
		num, _ := strconv.ParseFloat(numStr, 64)
		numberMap[num] = *big.NewInt(1)
	}

	for i := 0; i < 25; i++ {
		numberMap = GenerateNewList(numberMap)
	}

	sum := big.NewInt(0)
	for _, value := range numberMap {
		sum.Add(sum, &value)
	}

	return sum.String()
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	numberMap := make(map[float64]big.Int)

	numbers := strings.Fields(input[0])
	for _, numStr := range numbers {
		num, _ := strconv.ParseFloat(numStr, 64)
		numberMap[num] = *big.NewInt(1)
	}

	for i := 0; i < 75; i++ {
		numberMap = GenerateNewList(numberMap)
	}

	sum := big.NewInt(0)
	for _, value := range numberMap {
		sum.Add(sum, &value)
	}

	return sum.String()
}

func GenerateNewList(inputList map[float64]big.Int) map[float64]big.Int {
	newList := make(map[float64]big.Int)

	for key, value := range inputList {
		if key == 0 {
			if _, exists := newList[1]; !exists {
				temp := new(big.Int).Set(&value)
				newList[1] = *temp
			} else {
				temp := newList[1]
				temp.Add(&temp, &value)
				newList[1] = temp
			}
			continue
		}

		keyStr := fmt.Sprintf("%f", key)
		keyStr = strings.TrimRight(strings.TrimRight(keyStr, "0"), ".")

		if len(keyStr)%2 == 0 {
			halfLen := len(keyStr) / 2
			firstHalf, _ := strconv.ParseFloat(keyStr[:halfLen], 64)
			secondHalf, _ := strconv.ParseFloat(keyStr[halfLen:halfLen*2], 64)

			if _, exists := newList[firstHalf]; !exists {
				temp := new(big.Int).Set(&value)
				newList[firstHalf] = *temp
			} else {
				temp := newList[firstHalf]
				temp.Add(&temp, &value)
				newList[firstHalf] = temp
			}

			if _, exists := newList[secondHalf]; !exists {
				temp := new(big.Int).Set(&value)
				newList[secondHalf] = *temp
			} else {
				temp := newList[secondHalf]
				temp.Add(&temp, &value)
				newList[secondHalf] = temp
			}

			continue
		}

		newKey := key * 2024
		if _, exists := newList[newKey]; !exists {
			temp := new(big.Int).Set(&value)
			newList[newKey] = *temp
		} else {
			temp := newList[newKey]
			temp.Add(&temp, &value)
			newList[newKey] = temp
		}
	}

	return newList
}
