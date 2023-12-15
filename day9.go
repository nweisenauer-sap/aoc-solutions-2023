package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type day9 struct {
	input string
	oasis []series
}

func (d *day9) parseInput() {
	rawInput, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	input := string(rawInput)
	inputSlice := strings.Split(input, "\n")
	for _, line := range inputSlice {
		series := series{
			initial:     []int{},
			predictions: []int{},
		}
		lineSlice := strings.Fields(line)
		for _, num := range lineSlice {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			series.initial = append(series.initial, n)
		}
		d.oasis = append(d.oasis, series)
	}
}

func (d *day9) predictAll() {
	for i, series := range d.oasis {
		series.predict()
		d.oasis[i] = series
	}
}

type series struct {
	initial     []int
	predictions []int
	beginnings  []int
	next        int
	previous    int
}

func (s *series) predict() {
	s.predictRecursive(s.initial)
	s.predictions = append(s.predictions, s.initial[len(s.initial)-1])
	s.beginnings = append(s.beginnings, s.initial[0])

	next := 0
	for _, num := range s.predictions {
		next += num
	}
	s.next = next

	previous := 0
	for _, num := range s.beginnings {
		previous = num - previous
	}
	s.previous = previous
}

func (s *series) predictRecursive(series []int) {
	l := len(series)
	prediction := make([]int, l-1)
	for i := 0; i < l-1; i++ {
		prediction[i] = series[i+1] - series[i]
	}
	if l > 2 && s.predictionIsNotAllZeroes(prediction) {
		s.predictRecursive(prediction)
	}
	s.predictions = append(s.predictions, prediction[l-2])
	s.beginnings = append(s.beginnings, prediction[0])
}

func (s *series) predictionIsNotAllZeroes(prediction []int) bool {
	for _, num := range prediction {
		if num != 0 {
			return true
		}
	}
	return false
}

func day9Part1() string {
	d := day9{
		input: "input/day9_input.txt",
		oasis: []series{},
	}
	d.parseInput()
	d.predictAll()
	sum := 0
	for _, series := range d.oasis {
		sum += series.next
	}
	return fmt.Sprintf("%d", sum)
}

func day9Part2() string {
	d := day9{
		input: "input/day9_input.txt",
		oasis: []series{},
	}
	d.parseInput()
	d.predictAll()
	sum := 0
	for _, series := range d.oasis {
		sum += series.previous
	}
	return fmt.Sprintf("%d", sum)
}
