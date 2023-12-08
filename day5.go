package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type day5 struct {
	input   string
	almanac Almanac
}

func (d *day5) parseInput(seedRanges bool) {
	rawInput, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	input := string(rawInput)
	inputSlice := strings.Split(input, "\n")
	mapStart := 0
	for i, line := range inputSlice {
		switch {
		case i == 0 && !seedRanges:
			d.parseSeeds(line)
		case i == 0 && seedRanges:
			d.parseSeedRanges(line)
		case len(line) == 0 && mapStart == 0:
			mapStart = i + 1
		case len(line) == 0 && mapStart != 0:
			d.parseMapping(inputSlice[mapStart:i])
			mapStart = i + 1
		case i == len(inputSlice)-1:
			d.parseMapping(inputSlice[mapStart:])
		default:
			continue
		}
	}
}

func (d *day5) parseSeeds(line string) {
	line, found := strings.CutPrefix(line, "seeds: ")
	if !found {
		panic("seeds not found")
	}
	seedsSlice := strings.Split(line, " ")
	for _, seed := range seedsSlice {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		d.almanac.seeds[seedInt] = 1
	}
}

func (d *day5) parseSeedRanges(line string) {
	line, found := strings.CutPrefix(line, "seeds: ")
	if !found {
		panic("seeds not found")
	}
	seedRangesSlice := strings.Split(line, " ")
	startingSeed := -1
	for _, seed := range seedRangesSlice {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		if startingSeed == -1 {
			startingSeed = seedInt
		} else {
			d.almanac.seeds[startingSeed] = seedInt
			startingSeed = -1
		}
	}
}

func (d *day5) parseMapping(line []string) {
	m := Mapping{
		ranges: []Range{},
	}
	m.name = strings.Split(line[0], " ")[0]
	for _, line := range line[1:] {
		r := Range{}
		numbers := strings.Split(line, " ")
		var err error
		r.destination, err = strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		r.source, err = strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		r.length, err = strconv.Atoi(numbers[2])
		if err != nil {
			panic(err)
		}
		m.ranges = append(m.ranges, r)
	}
	d.almanac.mappings = append(d.almanac.mappings, m)
}

type Almanac struct {
	seeds    map[int]int
	mappings []Mapping
}

func (a *Almanac) findLowestFinalMapValue() int {
	lowestValue := math.MaxInt64
	lowestValues := make(chan int, len(a.seeds))
	expectedResultCount := 0
	segmentSize := 1000000
	for startSeed, rangeSize := range a.seeds {
		if rangeSize >= segmentSize {
			for i := 0; i < rangeSize; i += segmentSize {
				expectedResultCount++
				go func(startSeed int, rangeSize int) {
					lowestValues <- a.findLowestFinalMapValueForSegment(startSeed, rangeSize)
				}(startSeed+i, segmentSize)
			}
		}
		remainder := rangeSize % segmentSize
		if remainder > 0 {
			expectedResultCount++
			go func(startSeed int, rangeSize int) {
				lowestValues <- a.findLowestFinalMapValueForSegment(startSeed, rangeSize)
			}(startSeed+rangeSize-remainder, remainder)
		}
	}
	for i := 0; i < expectedResultCount; i++ {
		nextLowestValue := <-lowestValues
		if nextLowestValue < lowestValue {
			lowestValue = nextLowestValue
		}
	}

	return lowestValue
}

func (a *Almanac) findLowestFinalMapValueForSegment(startSeed int, rangeSize int) int {
	lowestValue := math.MaxInt64
	for seed := startSeed; seed < startSeed+rangeSize; seed++ {
		nextMapKey := seed
		for _, mapping := range a.mappings {
			nextMapKey = mapping.trueMap(nextMapKey)
		}
		if nextMapKey < lowestValue {
			lowestValue = nextMapKey
		}
	}
	return lowestValue
}

type Mapping struct {
	name   string
	ranges []Range
}

func (m *Mapping) trueMap(x int) int {
	for _, r := range m.ranges {
		if r.isInRange(x) {
			return r.destination + (x - r.source)
		}
	}
	return x
}

type Range struct {
	destination int
	source      int
	length      int
}

func (r *Range) isInRange(x int) bool {
	return x >= r.source && x < r.source+r.length
}

func day5Part1() string {
	d := day5{
		input: "input/day5_input.txt",
		almanac: Almanac{
			seeds:    map[int]int{},
			mappings: []Mapping{},
		},
	}
	d.parseInput(false)
	result := d.almanac.findLowestFinalMapValue()
	return strconv.Itoa(result)
}

func day5Part2() string {
	fmt.Println("Please be patient, this takes up to a minute to run ...")
	defer func(start time.Time) {
		fmt.Println("Computation Time:", time.Since(start))
	}(time.Now())
	d := day5{
		input: "input/day5_input.txt",
		almanac: Almanac{
			seeds:    map[int]int{},
			mappings: []Mapping{},
		},
	}
	d.parseInput(true)
	result := d.almanac.findLowestFinalMapValue()
	return strconv.Itoa(result)
}
