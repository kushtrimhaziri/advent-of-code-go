package day_13

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type EquationValues struct {
	A1, B1, A2, B2, C1, C2 int64
}

type Solution struct {
	X, Y int64
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

// processBlocks handles the common logic for both parts
func processBlocks(input []string, amplifier int64) int64 {
	var result int64
	blockSize := 3

	for i := 0; i < len(input); i += blockSize + 1 {
		block := strings.Join(input[i:i+blockSize], "\n")
		values, err := parseEquationValues(block, amplifier)
		if err != nil {
			fmt.Println("Error parsing input block:", err)
			continue
		}

		if solution := solveLinearEquations(values); solution != nil {
			result += solution.X*3 + solution.Y
		}
	}
	return result
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	return strconv.FormatInt(processBlocks(input, 0), 10)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return strconv.FormatInt(processBlocks(input, 10000000000000), 10)
}

func solveLinearEquations(values EquationValues) *Solution {
	determinant := values.A1*values.B2 - values.A2*values.B1
	if determinant == 0 {
		return nil
	}

	xNumerator := values.C1*values.B2 - values.C2*values.B1
	yNumerator := values.A1*values.C2 - values.A2*values.C1

	if xNumerator%determinant != 0 || yNumerator%determinant != 0 {
		return nil
	}

	x := xNumerator / determinant
	y := yNumerator / determinant

	if x >= 0 && y >= 0 {
		return &Solution{X: x, Y: y}
	}

	return nil
}

func parseEquationValues(input string, amplifier int64) (EquationValues, error) {
	pattern := "[+-]?\\d+"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)

	if len(matches) != 6 {
		return EquationValues{}, fmt.Errorf("expected exactly 6 numbers, found %d", len(matches))
	}

	a1, _ := strconv.ParseInt(matches[0], 10, 64)
	a2, _ := strconv.ParseInt(matches[1], 10, 64)
	b1, _ := strconv.ParseInt(matches[2], 10, 64)
	b2, _ := strconv.ParseInt(matches[3], 10, 64)
	c1, _ := strconv.ParseInt(matches[4], 10, 64)
	c2, _ := strconv.ParseInt(matches[5], 10, 64)

	c1 += amplifier
	c2 += amplifier

	return EquationValues{A1: a1, B1: b1, A2: a2, B2: b2, C1: c1, C2: c2}, nil
}
