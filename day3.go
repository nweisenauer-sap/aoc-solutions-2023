package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type day3 struct {
	input string
	gears map[coordinate]gearSlot
}

// create enum of "slot types" with possible values "empty","symbol" or "part"
const (
	empty = iota
	symbol
	digit
	gear
)

type gearSlot struct {
	neighbors []int
}

func (g gearSlot) ratio() int {
	if len(g.neighbors) != 2 {
		return 0
	}
	return g.neighbors[0] * g.neighbors[1]
}

type coordinate struct {
	row int
	col int
}

type slot struct {
	stype int
	value string
}

type engine struct {
	rows [][]slot
}

func (d *day3) parseEngineSchematic() engine {
	e := engine{}

	rawSchematic, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	schematic := string(rawSchematic)
	lines := strings.Split(schematic, "\n")

	for i, line := range lines {
		row := []slot{}
		for j, char := range line {
			switch {
			case unicode.IsNumber(char):
				row = append(row, slot{stype: digit, value: string(char)})
			case char == '*':
				row = append(row, slot{stype: gear})
				d.gears[coordinate{row: i, col: j}] = gearSlot{neighbors: []int{}}
			case char == '.':
				row = append(row, slot{stype: empty})
			default:
				row = append(row, slot{stype: symbol})
			}
		}
		e.rows = append(e.rows, row)
	}
	return e
}

func (d *day3) analyzeEngineSchematic(e engine) int {
	parts := []string{}
	for i, row := range e.rows {
		numStr := ""
		for j, slot := range row {
			if slot.stype == digit {
				numStr += slot.value
				if j < len(row)-1 {
					continue
				}
			}
			if numStr != "" {
				if d.numIsPart(e, numStr, i, j) {
					parts = append(parts, numStr)
				}
				numStr = ""
			}
		}
	}
	return d.sumAllParts(parts)
}

func (d *day3) numIsPart(e engine, numStr string, row int, lastCol int) bool {
	row1 := row - 1
	row2 := row + 1
	col1 := lastCol - len(numStr) - 1
	col2 := lastCol

	isPart := false

	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			if i == row && j > col1 && j < col2 {
				continue
			}
			if i < 0 || j < 0 || i >= len(e.rows) || j >= len(e.rows[i]) {
				continue
			}
			if e.rows[i][j].stype == symbol {
				isPart = true
			}
			if e.rows[i][j].stype == gear {
				isPart = true
				oldNeighbours := d.gears[coordinate{row: i, col: j}].neighbors
				number, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				d.gears[coordinate{row: i, col: j}] = gearSlot{neighbors: append(oldNeighbours, number)}
			}
		}
	}
	return isPart
}

func (d *day3) sumAllParts(parts []string) int {
	sum := 0
	for _, part := range parts {
		p, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		sum += p
	}
	return sum
}

func day3Part1() string {
	d := day3{
		input: "input/day3_input.txt",
		gears: map[coordinate]gearSlot{},
	}
	e := d.parseEngineSchematic()
	s := d.analyzeEngineSchematic(e)
	return fmt.Sprintf("%v", s)
}

func day3Part2() string {
	d := day3{
		input: "input/day3_input.txt",
		gears: map[coordinate]gearSlot{},
	}
	e := d.parseEngineSchematic()
	d.analyzeEngineSchematic(e)
	s := 0
	for _, gear := range d.gears {
		s += gear.ratio()
	}
	return fmt.Sprintf("%v", s)
}
