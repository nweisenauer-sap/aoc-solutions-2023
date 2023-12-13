package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

const joker = 11

type day7 struct {
	input string
	hands []Hand
}

func (d *day7) parseInput(useJokers bool) {
	rawInput, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	input := string(rawInput)
	inputSlice := strings.Split(input, "\n")

	for _, line := range inputSlice {
		if len(line) == 0 {
			continue
		}
		hand := Hand{
			cards: [5]Card{},
			bid:   0,
		}
		lineSlice := strings.Split(line, " ")

		cards := lineSlice[0]
		bid, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			panic(err)
		}
		hand.bid = bid

		for i, card := range cards {
			cardInt := d.cardToInt(card)
			hand.cards[i] = cardInt
		}
		hand.typeValue = d.calcTypeValue(hand, useJokers)

		if useJokers {
			for i, card := range hand.cards {
				if card == joker {
					hand.cards[i] = 0
				}
			}
		}

		d.hands = append(d.hands, hand)
	}
}

func (d *day7) calcTypeValue(h Hand, useJokers bool) int {
	buckets := make(map[int]int)
	for _, card := range h.cards {
		buckets[card]++
	}

	largestBucket := 0
	secondLargestBucket := 0
	jokers := buckets[joker]

	for i, bucket := range buckets {
		if i == joker && useJokers {
			continue
		}
		if bucket >= largestBucket {
			secondLargestBucket = largestBucket
			largestBucket = bucket
		} else if bucket >= secondLargestBucket {
			secondLargestBucket = bucket
		}
	}

	if useJokers {
		largestBucket += jokers
	}

	switch {
	case largestBucket == 5:
		return 6
	case largestBucket == 4:
		return 5
	case largestBucket == 3 && secondLargestBucket == 2:
		return 4
	case largestBucket == 3:
		return 3
	case largestBucket == 2 && secondLargestBucket == 2:
		return 2
	case largestBucket == 2:
		return 1
	default:
		return 0
	}
}

func (d *day7) cardToInt(card rune) Card {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func (d *day7) sortHands() {
	slices.SortFunc(d.hands, func(h1, h2 Hand) int {
		return h1.Compare(&h2)
	})
}

func (d *day7) totalWinnings() int {
	d.sortHands()
	total := 0
	for i, hand := range d.hands {
		rank := i + 1
		total += (hand.bid * rank)
	}
	return total
}

type Card = int

type Hand struct {
	cards     [5]Card
	bid       int
	typeValue int
}

func (h *Hand) Compare(h2 *Hand) int {
	if h.typeValue != h2.typeValue {
		return h.typeValue - h2.typeValue
	}
	for i := 0; i < 5; i++ {
		if h.cards[i] != h2.cards[i] {
			return h.cards[i] - h2.cards[i]
		}
	}
	return 0
}

func day7Part1() string {
	d := day7{
		input: "input/day7_input.txt",
		hands: []Hand{},
	}
	d.parseInput(false)
	return strconv.Itoa(d.totalWinnings())
}

func day7Part2() string {
	d := day7{
		input: "input/day7_input.txt",
		hands: []Hand{},
	}
	d.parseInput(true)
	return strconv.Itoa(d.totalWinnings())
}
