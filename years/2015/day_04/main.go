package day_04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func findHashWithPrefix(input string, zeros int) string {
	prefix := strings.Repeat("0", zeros)
	var numberToBeAdded = 0

	for {
		hashOfInput := GetMD5Hash(input + strconv.Itoa(numberToBeAdded))
		if strings.HasPrefix(hashOfInput, prefix) {
			break
		}
		numberToBeAdded++
	}
	return strconv.Itoa(numberToBeAdded)
}

// Part1 solves the first part of the exercise
func Part1(input string) string {
	return findHashWithPrefix(input, 5)
}

// Part2 solves the second part of the exercise
func Part2(input string) string {
	return findHashWithPrefix(input, 6)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
