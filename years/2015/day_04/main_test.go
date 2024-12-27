package day_04_test

import (
	"github.com/kushtrimhaziri/advent_of_code_go/internal"
	"github.com/kushtrimhaziri/advent_of_code_go/years/2015/day_04"
	"github.com/kushtrimhaziri/advent_of_code_go/years/2023/day_0"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	input := internal.LoadFirstInputLine("input_1_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_1.txt")
	result := day_04.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_2_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_2.txt")
	result := day_0.Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}
