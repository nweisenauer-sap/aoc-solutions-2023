package main

import (
	"os"
	"strconv"
	"strings"
)

type day1 struct {
}

func (d *day1) parseCalibration(input string) string {
	calibration := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// find the first single digit in the line
		firstIndex := strings.IndexFunc(line, func(r rune) bool { return r >= '0' && r <= '9' })
		firstDigit := string(line[firstIndex])
		// find the last single digit in the line
		lastIndex := strings.LastIndexFunc(line, func(r rune) bool { return r >= '0' && r <= '9' })
		lastDigit := string(line[lastIndex])
		// create a 2-digit number from first and last
		number, _ := strconv.Atoi(firstDigit + lastDigit)
		// add the sum to the calibration
		calibration += number
	}

	return strconv.Itoa(calibration)
}

func day1Part1() string {
	d := day1{}
	content, err := os.ReadFile("input/day1_input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	return d.parseCalibration(input)
}

func day1Part2() string {
	d := day1{}
	mapWordNumbersToSlightlyLessWordyNumbers := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	content, err := os.ReadFile("input/day1_input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	// find the words and replace them with numbers
	for word, number := range mapWordNumbersToSlightlyLessWordyNumbers {
		input = strings.ReplaceAll(input, word, number)
	}

	return d.parseCalibration(input)
}
