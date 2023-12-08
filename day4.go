package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

type day4 struct {
	input        string
	scratchCards []scratchCard
}

type scratchCard struct {
	winningNumbers []int
	ownNumbers     []int
	matches        int
	points         int
	instances      int
}

func (s *scratchCard) scoreCard() int {
	for _, ownNumber := range s.ownNumbers {
		if slices.Contains(s.winningNumbers, ownNumber) {
			s.scoreNumber()
		}
	}
	return s.points
}

func (s *scratchCard) scoreNumber() {
	s.matches++
	if s.points == 0 {
		s.points = 1
	} else {
		s.points *= 2
	}
}

func (d *day4) parseInput() {
	rawCards, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	cards := string(rawCards)
	cardSlice := strings.Split(cards, "\n")
	for _, card := range cardSlice {
		scratchCard := scratchCard{
			points:    0,
			matches:   0,
			instances: 1,
		}
		scratchCard.winningNumbers = d.parseWinningNumbers(card)
		scratchCard.ownNumbers = d.parseOwnNumbers(card)
		d.scratchCards = append(d.scratchCards, scratchCard)
	}
}

func (d *day4) parseWinningNumbers(card string) []int {
	result := []int{}
	card = card[strings.Index(card, ":")+1:]
	card = card[:strings.Index(card, "|")]
	numbersSlice := strings.Split(card, " ")
	for _, number := range numbersSlice {
		if number == "" || number == " " {
			continue
		}
		trimmedNumber, error := strconv.Atoi(strings.TrimSpace(number))
		if error != nil {
			panic(error)
		}
		result = append(result, trimmedNumber)
	}
	return result
}

func (d *day4) parseOwnNumbers(card string) []int {
	result := []int{}
	card = card[strings.Index(card, "|")+1:]
	numbersSlice := strings.Split(card, " ")
	for _, number := range numbersSlice {
		if number == "" || number == " " {
			continue
		}
		trimmedNumber, error := strconv.Atoi(strings.TrimSpace(number))
		if error != nil {
			panic(error)
		}
		result = append(result, trimmedNumber)
	}
	return result
}

func (d *day4) totalScore() int {
	totalScore := 0
	for i, card := range d.scratchCards {
		totalScore += card.scoreCard()
		d.scratchCards[i] = card
	}
	return totalScore
}

func (d *day4) calculateInstances() {
	for i, card := range d.scratchCards {
		firstSuccessor := i + 1
		lastSuccessor := i + card.matches
		for j := firstSuccessor; j <= lastSuccessor && j < len(d.scratchCards); j++ {
			d.scratchCards[j].instances += card.instances
		}
	}
}

func (d *day4) totalInstances() int {
	totalInstances := 0
	for _, card := range d.scratchCards {
		totalInstances += card.instances
	}
	return totalInstances
}

func day4Part1() string {
	d := day4{
		input:        "input/day4_input.txt",
		scratchCards: []scratchCard{},
	}
	d.parseInput()
	return strconv.Itoa(d.totalScore())
}

func day4Part2() string {
	d := day4{
		input:        "input/day4_input.txt",
		scratchCards: []scratchCard{},
	}
	d.parseInput()
	d.totalScore()
	d.calculateInstances()
	return strconv.Itoa(d.totalInstances())
}
