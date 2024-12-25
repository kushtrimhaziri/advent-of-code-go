package day_25

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

	locks, keys := processBlocks(input)

	countOfKeys := 0

	for _, lock := range locks {
		for _, key := range keys {
			itIsAPair := true
			for k := 0; k < len(lock); k++ {
				if lock[k]+key[k]+1 >= 7 {
					itIsAPair = false
					break
				}

			}
			if itIsAPair {
				countOfKeys++
			}
		}
	}

	return strconv.Itoa(countOfKeys)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

func processBlocks(input []string) (locks [][]int, keys [][]int) {
	var locksOfBlocks [][]int
	var keysOfBlocks [][]int
	blockSize := 7

	for i := 0; i < len(input); i += blockSize + 1 {
		tracker := make([]int, 5)
		for i := range tracker {
			tracker[i] = -1
		}
		println(i)
		for j := i; j < i+blockSize; j++ {
			for k := 0; k < len(input[i]); k++ {

				if input[j][k] == '#' {
					tracker[k]++
				}
			}
		}

		if input[i][0] == '#' {
			locksOfBlocks = append(locksOfBlocks, tracker)
		} else {
			keysOfBlocks = append(keysOfBlocks, tracker)

		}
	}

	return locksOfBlocks, keysOfBlocks
}
