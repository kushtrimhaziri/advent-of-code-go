package day_04

import (
	"fmt"
)

type Point struct {
	Row, Col int
}

type LinePoints struct {
	Start, End Point
}

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
	xmasWord := "XMAS"
	occurrences := 0
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != 'X' {
				continue
			}

			for _, dir := range directions {
				newRow := i + dir[0]
				newCol := j + dir[1]

				if newRow >= 0 && newRow < len(input) &&
					newCol >= 0 && newCol < len(input[i]) &&
					input[newRow][newCol] == 'M' {

					found := true
					for idx, char := range xmasWord[1:] {

						checkRow := i + dir[0]*(idx+1)
						checkCol := j + dir[1]*(idx+1)

						if checkRow < 0 || checkRow >= len(input) ||
							checkCol < 0 || checkCol >= len(input[i]) ||
							input[checkRow][checkCol] != byte(char) {
							found = false
							break
						}
					}
					if found {
						occurrences++
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", occurrences)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	diagonalMAS := make(map[Point][]LinePoints)
	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'A' {
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(input) && j+1 < len(input[i]) {
					if (input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') ||
						(input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M') {
						center := Point{Row: i, Col: j}
						pattern := LinePoints{
							Start: Point{Row: i - 1, Col: j - 1},
							End:   Point{Row: i + 1, Col: j + 1},
						}
						diagonalMAS[center] = append(diagonalMAS[center], pattern)
					}

					if (input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S') ||
						(input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M') {
						center := Point{Row: i, Col: j}
						pattern := LinePoints{
							Start: Point{Row: i - 1, Col: j + 1},
							End:   Point{Row: i + 1, Col: j - 1},
						}
						diagonalMAS[center] = append(diagonalMAS[center], pattern)
					}
				}
			}
		}
	}

	for _, patterns := range diagonalMAS {
		// Create a map to check for unique patterns
		uniquePatterns := make(map[LinePoints]bool)
		for _, pattern := range patterns {
			uniquePatterns[pattern] = true
		}
		if len(uniquePatterns) > 1 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}
