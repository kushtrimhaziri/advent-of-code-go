package day_03

import (
	"fmt"
	"strconv"
)

type Coord struct {
	x, y int
}

// Run function of the daily challenge
func Run(input string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func updateCoord(coord Coord, direction rune) Coord {
	switch direction {
	case '>':
		coord.x++
	case '<':
		coord.x--
	case '^':
		coord.y++
	case 'v':
		coord.y--
	}
	return coord
}

// Part1 solves the first part of the exercise
func Part1(input string) string {
	houses := make(map[Coord]int)
	currentCoord := Coord{}
	houses[currentCoord] = 1

	for _, direction := range input {
		currentCoord = updateCoord(currentCoord, direction)
		houses[currentCoord]++
	}
	return strconv.Itoa(len(houses))
}

// Part2 solves the second part of the exercise
func Part2(input string) string {
	houses := make(map[Coord]int)
	coords := [2]Coord{{}, {}}
	houses[coords[0]] = 1

	for i, direction := range input {
		current := i % 2
		coords[current] = updateCoord(coords[current], direction)
		houses[coords[current]]++
	}
	return strconv.Itoa(len(houses))
}
