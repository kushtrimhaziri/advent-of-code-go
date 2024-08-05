package day_01

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type PointWithCoordinatesOnly struct {
	X, Y int
}

type Point struct {
	X, Y int
	Direction
}

type UpdatedPoint struct {
	Point  Point
	Exists bool
}

var directionChanges = map[Direction]map[string]Point{
	Up: {
		"R": {X: 1, Y: 0, Direction: Right},
		"L": {X: -1, Y: 0, Direction: Left},
	},
	Down: {
		"R": {X: -1, Y: 0, Direction: Left},
		"L": {X: 1, Y: 0, Direction: Right},
	},
	Left: {
		"R": {X: 0, Y: 1, Direction: Up},
		"L": {X: 0, Y: -1, Direction: Down},
	},
	Right: {
		"R": {X: 0, Y: -1, Direction: Down},
		"L": {X: 0, Y: 1, Direction: Up},
	},
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

// Part1 solves the first part of the exercise
func Part1(input string) string {
	commands := parseCommands(input)
	point := Point{X: 0, Y: 0, Direction: Up}

	for _, cmd := range commands {
		point = updatePointForPart1(point, point.Direction, cmd.command, cmd.amount)
	}

	sumOfCommands := math.Abs(float64(point.X)) + math.Abs(float64(point.Y))

	return strconv.FormatFloat(sumOfCommands, 'f', -1, 64)
}

// Part2 solves the second part of the exercise
func Part2(input string) string {
	commands := parseCommands(input)
	point := Point{X: 0, Y: 0, Direction: Up}
	pointMap := make(map[PointWithCoordinatesOnly]bool)
	exists := false
	updatedPoint := UpdatedPoint{point, exists}

	for _, cmd := range commands {
		updatedPoint = updatePointForPart2(updatedPoint, updatedPoint.Point.Direction, cmd.command, cmd.amount, pointMap)

		if updatedPoint.Exists {
			sumOfCommands := math.Abs(float64(updatedPoint.Point.X)) + math.Abs(float64(updatedPoint.Point.Y))

			return strconv.FormatFloat(sumOfCommands, 'f', -1, 64)
		}
	}
	return ""
}

func updatePointForPart1(currentPoint Point, direction Direction, command string, amount int) Point {
	change := directionChanges[direction][command]

	if change.X != 0 {
		currentPoint.X += change.X * amount
	} else {
		currentPoint.Y += change.Y * amount

	}
	currentPoint.Direction = change.Direction

	return currentPoint
}

func updatePointForPart2(currentPoint UpdatedPoint, direction Direction, command string, amount int, pointMap map[PointWithCoordinatesOnly]bool) UpdatedPoint {
	change := directionChanges[direction][command]

	if change.X != 0 {
		for i := 0; i < amount; i++ {
			currentPoint.Point.X += change.X * 1
			newPoint := PointWithCoordinatesOnly{X: currentPoint.Point.X, Y: currentPoint.Point.Y}
			if _, existsNewPoint := pointMap[newPoint]; existsNewPoint {
				currentPoint.Exists = true
				return currentPoint
			}
			pointMap[newPoint] = true
		}
	} else {
		for i := 0; i < amount; i++ {
			currentPoint.Point.Y += change.Y * 1
			newPoint := PointWithCoordinatesOnly{X: currentPoint.Point.X, Y: currentPoint.Point.Y}

			if _, existsNewPoint := pointMap[newPoint]; existsNewPoint {
				currentPoint.Exists = true
				return currentPoint
			}
			pointMap[newPoint] = true
		}
	}

	currentPoint.Point.Direction = change.Direction
	return currentPoint
}
func parseCommands(commandsStr string) []struct {
	command string
	amount  int
} {
	commandSegments := strings.Split(commandsStr, ", ")

	parsedCommands := make([]struct {
		command string
		amount  int
	}, len(commandSegments))

	for i, segment := range commandSegments {
		command := segment[:1]

		amount, err := strconv.Atoi(segment[1:])
		if err != nil {
			fmt.Printf("Error parsing amount from %s: %v\n", segment, err)
			continue
		}

		parsedCommands[i] = struct {
			command string
			amount  int
		}{command, amount}
	}

	return parsedCommands
}
