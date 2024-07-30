package day_01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var Numbers map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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

	result := 0
	for _, value := range input {
		digits := findDigits(value)

		numberToBeAdded := ""
		if len(digits) >= 2 {
			numberToBeAdded = digits[0] + digits[len(digits)-1]
		} else {
			numberToBeAdded = digits[0] + digits[0]
		}

		parsedInt, err := strconv.Atoi(numberToBeAdded)
		if err != nil {
			fmt.Println("Error parsing result:", err)
			return ""
		}
		result += parsedInt
	}
	finalResult := strconv.Itoa(result)

	return finalResult
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {

	result := 0
	for _, value := range input {
		var numbers []int
		var builder strings.Builder
		for i := 0; i < len(value); i++ {
			asciiCode := value[i]
			builder.WriteString(string(asciiCode))
			if asciiCode >= '0' && asciiCode <= '9' {
				numbers = append(numbers, int(asciiCode-'0'))
				builder = strings.Builder{}
				continue
			}

			if builder.Len() < 3 {
				continue
			}
			numberObtained, isValid := parseNumberString(&builder)

			if isValid {
				parsedInt := Numbers[numberObtained]
				numbers = append(numbers, parsedInt)
			}

			if builder.Len() > 5 {
				removeFirstCharacter(&builder)
				continue
			}
		}
		numberToBeAdded := 0
		if len(numbers) >= 2 {
			numberToBeAdded = numbers[0]*10 + numbers[len(numbers)-1]
		} else {
			numberToBeAdded = numbers[0]*10 + numbers[0]
		}

		result += numberToBeAdded
	}

	finalResult := strconv.Itoa(result)
	return finalResult
}

// Helper functions
func findDigits(input string) []string {
	re := regexp.MustCompile(`\d`)

	matches := re.FindAllString(input, -1)

	return matches
}

func parseNumberString(input *strings.Builder) (string, bool) {
	testingStringBuilder := strings.Builder{}
	testingStringBuilder.WriteString(input.String())

	for testingStringBuilder.Len() >= 3 {
		if _, ok := Numbers[testingStringBuilder.String()]; ok {
			return testingStringBuilder.String(), true
		}

		currentTestingStringBuilder := testingStringBuilder.String()[1:]
		testingStringBuilder.Reset()
		testingStringBuilder.WriteString(currentTestingStringBuilder)
	}

	return "", false
}

func removeFirstCharacter(builder *strings.Builder) {
	currentContent := builder.String()

	if len(currentContent) > 0 {
		newContent := currentContent[1:]

		builder.Reset()
		builder.WriteString(newContent)
	}
}
