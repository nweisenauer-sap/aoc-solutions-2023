package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type day6 struct {
	input string
	races []Race
}

func (d *day6) parseInput(longRace bool) {
	rawInput, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	input := string(rawInput)
	inputSlice := strings.Split(input, "\n")

	firstLine, found := strings.CutPrefix(inputSlice[0], "Time: ")
	if !found {
		panic("Time not found")
	}
	secondLine, found := strings.CutPrefix(inputSlice[1], "Distance: ")
	if !found {
		panic("Distance not found")
	}

	if longRace {
		firstLine = strings.ReplaceAll(firstLine, " ", "")
		secondLine = strings.ReplaceAll(secondLine, " ", "")

		timeInt, err := strconv.Atoi(firstLine)
		if err != nil {
			panic(err)
		}
		distanceInt, err := strconv.Atoi(secondLine)
		if err != nil {
			panic(err)
		}

		race := Race{
			time:          timeInt,
			distance:      distanceInt,
			winningSpeeds: []int{},
		}
		d.races = append(d.races, race)
		return
	}

	timesSlice := strings.Split(firstLine, " ")
	distancesSlice := strings.Split(secondLine, " ")

	timesSlice = d.deleteEmptyStrings(timesSlice)
	distancesSlice = d.deleteEmptyStrings(distancesSlice)

	for i, time := range timesSlice {
		distance := distancesSlice[i]
		timeInt, err := strconv.Atoi(time)
		if err != nil {
			panic(err)
		}
		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			panic(err)
		}
		race := Race{
			time:          timeInt,
			distance:      distanceInt,
			winningSpeeds: []int{},
		}
		d.races = append(d.races, race)
	}
}

func (d *day6) deleteEmptyStrings(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}

func (d *day6) findAllWinningSpeeds() {
	for i, race := range d.races {
		speed, _ := race.findLowestWinningSpeed()
		race.winningSpeeds = append(race.winningSpeeds, speed)

		for speed++; race.checkWin(speed) <= 0; speed++ {
			race.winningSpeeds = append(race.winningSpeeds, speed)
		}
		d.races[i] = race
	}
}

type Race struct {
	time          int
	distance      int
	winningSpeeds []int
}

func (r *Race) findLowestWinningSpeed() (int, bool) {
	return sort.Find(r.time+1, func(i int) int {
		return r.checkWin(i)
	})
}

func (r *Race) checkWin(speed int) int {
	travelTime := r.time - speed
	return (r.distance + 1) - speed*travelTime
}

func day6Part1() string {
	d := day6{
		input: "input/day6_input.txt",
		races: []Race{},
	}
	d.parseInput(false)
	d.findAllWinningSpeeds()
	result := 1
	for _, r := range d.races {
		result *= len(r.winningSpeeds)
	}
	return strconv.Itoa(result)
}

func day6Part2() string {
	d := day6{
		input: "input/day6_input.txt",
		races: []Race{},
	}
	d.parseInput(true)
	d.findAllWinningSpeeds()
	fmt.Printf("hehe: %q", "d.races[0].winningSpeeds")
	result := 1
	for _, r := range d.races {
		result *= len(r.winningSpeeds)
	}
	return strconv.Itoa(result)
}
