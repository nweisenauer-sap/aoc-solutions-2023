package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type day2 struct {
}

func (d *day2) parseInput() []game {
	content, err := os.ReadFile("input/day2_input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	games := strings.Split(input, "\n")

	var parsedGames []game
	for i, gameRaw := range games {

		gameRaw = strings.TrimPrefix(gameRaw, fmt.Sprintf("Game %v: ", i+1))
		drawsRaw := strings.Split(gameRaw, ";")

		var parsedDraws []draw
		for _, drawRaw := range drawsRaw {

			drawRaw = strings.TrimSpace(drawRaw)
			ballsRaw := strings.Split(drawRaw, ", ")

			var parsedDraw draw
			for _, ballRaw := range ballsRaw {

				ball := strings.Split(ballRaw, " ")
				count, _ := strconv.Atoi(ball[0])

				switch ball[1] {
				case "blue":
					parsedDraw.blue = count
				case "green":
					parsedDraw.green = count
				case "red":
					parsedDraw.red = count
				}
			}

			parsedDraws = append(parsedDraws, parsedDraw)
		}

		parsedGames = append(parsedGames, game{i + 1, parsedDraws})
	}
	return parsedGames
}

type bag struct {
	blue  int
	green int
	red   int
}

type draw struct {
	blue  int
	green int
	red   int
}

func (d draw) String() string {
	return fmt.Sprintf(" %v blue, %v green, %v red;", d.blue, d.green, d.red)
}

func (d draw) isValid(b bag) bool {
	return d.blue <= b.blue && d.green <= b.green && d.red <= b.red
}

type game struct {
	id    int
	draws []draw
}

func (g game) String() string {
	var output string = fmt.Sprintf("Game %v:", g.id)
	for _, draw := range g.draws {
		output += fmt.Sprintf("%v", draw)
	}
	output = strings.TrimSuffix(output, ";")
	output += "\n"
	return output
}

func (g game) isValid(b bag) bool {
	for _, draw := range g.draws {
		if !draw.isValid(b) {
			return false
		}
	}
	return true
}

func (g game) minCubesInBag() bag {
	maxBlue := 0
	maxGreen := 0
	maxRed := 0
	for _, draw := range g.draws {
		if draw.blue > maxBlue {
			maxBlue = draw.blue
		}
		if draw.green > maxGreen {
			maxGreen = draw.green
		}
		if draw.red > maxRed {
			maxRed = draw.red
		}
	}
	return bag{blue: maxBlue, green: maxGreen, red: maxRed}
}

func day2Part1() string {
	d := day2{}
	games := d.parseInput()
	// sum of all valid game IDs
	sum := 0
	for _, game := range games {
		if game.isValid(bag{blue: 14, green: 13, red: 12}) {
			sum += game.id
		}
	}

	return fmt.Sprintf("%v", sum)
}

func day2Part2() string {
	d := day2{}
	games := d.parseInput()
	// sum of the power of minimal bags
	sumOfPower := 0

	for _, game := range games {
		minBag := game.minCubesInBag()
		sumOfPower += (minBag.blue * minBag.green * minBag.red)
	}
	return fmt.Sprintf("%v", sumOfPower)
}
