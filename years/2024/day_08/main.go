package day_08

import (
	"fmt"
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

type Point struct {
	row, col int
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	antennas := make(map[rune][]Point)
	antinodes := make(map[Point]bool)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				point := Point{i, j}
				char := rune(input[i][j])
				antennas[char] = append(antennas[char], point)
			}
		}
	}

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				rowDiff := points[i].row - points[j].row
				colDiff := points[i].col - points[j].col

				possibleAntinode := Point{
					row: points[i].row + rowDiff,
					col: points[i].col + colDiff,
				}
				addAntinodeIfValid(possibleAntinode, input, antinodes)

				possibleAntinode2 := Point{
					row: points[j].row + (points[j].row - points[i].row),
					col: points[j].col + (points[j].col - points[i].col),
				}
				addAntinodeIfValid(possibleAntinode2, input, antinodes)
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	antennas := make(map[rune][]Point)
	antinodes := make(map[Point]bool)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				point := Point{i, j}
				antinodes[point] = true

				char := rune(input[i][j])
				antennas[char] = append(antennas[char], point)
			}
		}
	}

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				rowDiff := points[i].row - points[j].row
				colDiff := points[i].col - points[j].col

				possibleAntinode := Point{
					row: points[i].row + rowDiff,
					col: points[i].col + colDiff,
				}

				for possibleAntinode.row >= 0 && possibleAntinode.row < len(input) &&
					possibleAntinode.col >= 0 && possibleAntinode.col < len(input[0]) {
					antinodes[possibleAntinode] = true
					possibleAntinode.row += rowDiff
					possibleAntinode.col += colDiff
				}

				possibleAntinode2 := Point{
					row: points[j].row + (points[j].row - points[i].row),
					col: points[j].col + (points[j].col - points[i].col),
				}

				for possibleAntinode2.row >= 0 && possibleAntinode2.row < len(input) &&
					possibleAntinode2.col >= 0 && possibleAntinode2.col < len(input[0]) {
					antinodes[possibleAntinode2] = true
					possibleAntinode2.row += (points[j].row - points[i].row)
					possibleAntinode2.col += (points[j].col - points[i].col)
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes))
}

func addAntinodeIfValid(point Point, input []string, antinodes map[Point]bool) {
	if point.row >= 0 && point.row < len(input) &&
		point.col >= 0 && point.col < len(input[0]) {
		antinodes[point] = true
	}
}
