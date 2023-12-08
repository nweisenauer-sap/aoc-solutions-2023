package main

import "fmt"

func main() {
	// ask for user input regarding the day then execute both parts of that day sequentially
	fmt.Println("Which day would you like to run?")
	var day int
	fmt.Scanln(&day)
	var resultA string
	var resultB string
	switch day {
	case 1:
		resultA = day1Part1()
		resultB = day1Part2()
	case 2:
		resultA = day2Part1()
		resultB = day2Part2()
	case 3:
		resultA = day3Part1()
		resultB = day3Part2()
	case 4:
		resultA = day4Part1()
		resultB = day4Part2()
	case 5:
		resultA = day5Part1()
		resultB = day5Part2()
	case 6:
		resultA = day6Part1()
		resultB = day6Part2()
	case 7:
		resultA = day7Part1()
		resultB = day7Part2()
	case 8:
		resultA = day8Part1()
		resultB = day8Part2()
	case 9:
		resultA = day9Part1()
		resultB = day9Part2()
	case 10:
		resultA = day10Part1()
		resultB = day10Part2()
	case 11:
		resultA = day11Part1()
		resultB = day11Part2()
	case 12:
		resultA = day12Part1()
		resultB = day12Part2()
	case 13:
		resultA = day13Part1()
		resultB = day13Part2()
	case 14:
		resultA = day14Part1()
		resultB = day14Part2()
	case 15:
		resultA = day15Part1()
		resultB = day15Part2()
	case 16:
		resultA = day16Part1()
		resultB = day16Part2()
	case 17:
		resultA = day17Part1()
		resultB = day17Part2()
	case 18:
		resultA = day18Part1()
		resultB = day18Part2()
	case 19:
		resultA = day19Part1()
		resultB = day19Part2()
	case 20:
		resultA = day20Part1()
		resultB = day20Part2()
	case 21:
		resultA = day21Part1()
		resultB = day21Part2()
	case 22:
		resultA = day22Part1()
		resultB = day22Part2()
	case 23:
		resultA = day23Part1()
		resultB = day23Part2()
	case 24:
		resultA = day24Part1()
		resultB = day24Part2()
	default:
		fmt.Println("Invalid day, choose a day between 1 and 24")
	}

	fmt.Println("")
	fmt.Println("#=#=*=*=*=*=*=*=*=*=*=*=*")
	fmt.Println("||| Advent of Code 2023 |")
	fmt.Println("#-#-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println("|||                     |")
	fmt.Printf("|||       DAY %v         |\n", day)
	fmt.Println("||.                     .")
	fmt.Println("||   Part 1:", resultA)
	fmt.Println("||   Part 2:", resultB)
	fmt.Println("||.                     .")
	fmt.Println("|||                      \\")
	fmt.Println("|||\\______________________\\")
	fmt.Println("#=#=*=*=*=*=*=*=*=*=*=*=*")
}
