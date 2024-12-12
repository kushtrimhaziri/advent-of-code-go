package day_10

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

	var matrix [][]int

	for _, row := range input {
		var matrixRow []int
		for _, char := range row {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Printf("Error converting character to integer: %v\n", err)
				continue
			}
			matrixRow = append(matrixRow, num)
		}
		matrix = append(matrix, matrixRow)
	}

	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 9 {
				println(count)
				count += countReachableTrails(matrix, i, j)
			}
		}
	}

	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	var matrix [][]int

	for _, row := range input {
		var matrixRow []int
		for _, char := range row {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Printf("Error converting character to integer: %v\n", err)
				continue
			}
			matrixRow = append(matrixRow, num)
		}
		matrix = append(matrix, matrixRow)
	}

	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 9 {
				println(count)
				count += countDifferentTrails(matrix, i, j)
			}
		}
	}

	return strconv.Itoa(count)
}

func countDifferentTrails(matrix [][]int, startX, startY int) int {
	visited := make(map[string]bool)
	return exploreTrails(matrix, startX, startY, 9, visited)
}

func exploreTrails(matrix [][]int, x, y, currentValue int, visited map[string]bool) int {
	if currentValue == 0 {
		return 1
	}

	pos := fmt.Sprintf("%d,%d", x, y)
	if visited[pos] {
		return 0
	}

	visited[pos] = true
	defer delete(visited, pos) // took it after googling for DFS

	totalPaths := 0
	nextValue := currentValue - 1

	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) {
			if matrix[nx][ny] == nextValue {
				totalPaths += exploreTrails(matrix, nx, ny, nextValue, visited)
			}
		}
	}

	return totalPaths
}

func countReachableTrails(matrix [][]int, startX, startY int) int {
	visited := make(map[string]bool)
	reachableZeros := make(map[string]bool)
	exploreTrailheads(matrix, startX, startY, 9, visited, reachableZeros)
	return len(reachableZeros)
}

func exploreTrailheads(matrix [][]int, x, y, currentValue int, visited map[string]bool, reachableZeros map[string]bool) {
	pos := fmt.Sprintf("%d,%d", x, y)
	if visited[pos] {
		return
	}

	visited[pos] = true
	defer delete(visited, pos) // took it after googling for DFS

	if matrix[x][y] == 0 {
		reachableZeros[pos] = true
		return
	}

	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	nextValue := currentValue - 1
	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) {
			if matrix[nx][ny] == nextValue {
				exploreTrailheads(matrix, nx, ny, nextValue, visited, reachableZeros)
			}
		}
	}
}
