package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type day8 struct {
	input        string
	instructions string
	graph        map[string]node
}

func (d *day8) parseInput() {
	rawInput, err := os.ReadFile(d.input)
	if err != nil {
		panic(err)
	}
	input := string(rawInput)
	inputSlice := strings.Split(input, "\n")
	d.instructions = inputSlice[0]
	d.parseGraph(inputSlice[2:])
}

func (d *day8) parseGraph(graphSlice []string) {
	for _, line := range graphSlice {
		node := node{}
		lineSlice := strings.Split(line, " = ")
		node.id = lineSlice[0]
		edgesSlice := strings.Split(lineSlice[1], ", ")

		node.leftEdge = strings.TrimPrefix(edgesSlice[0], "(")
		node.rightEdge = strings.TrimSuffix(edgesSlice[1], ")")
		d.graph[node.id] = node
	}
}

func (d *day8) traverseGraph() int {
	currentNode := "AAA"
	steps := 0
	maxIterations := 10000

	for i := 0; i < maxIterations; i++ {
		for _, instruction := range d.instructions {
			switch instruction {
			case 'L':
				currentNode = d.graph[currentNode].leftEdge
			case 'R':
				currentNode = d.graph[currentNode].rightEdge
			}
			steps++
			if currentNode == "ZZZ" {
				return steps
			}
		}
	}
	return 0
}

func (d *day8) traverseGraphLikeAGhost() int {
	currentNodes := []string{}
	for id := range d.graph {
		if strings.HasSuffix(id, "A") {
			currentNodes = append(currentNodes, id)
		}
	}

	goRoutineCount := len(currentNodes)
	solutionChan := make(chan int, goRoutineCount)
	wg := sync.WaitGroup{}

	for _, currentNode := range currentNodes {
		wg.Add(1)
		go d.performGOstlyTraversal(currentNode, solutionChan, &wg)
	}

	wg.Wait()
	close(solutionChan)
	solutions := []int{}
	for solution := range solutionChan {
		solutions = append(solutions, solution)
	}

	if len(solutions) > 1 {
		return d.LCM(solutions[0], solutions[1], solutions[2:]...)
	}

	return 0
}

func (d *day8) performGOstlyTraversal(node string, solutionChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	steps := 0
	for {
		for _, instruction := range d.instructions {
			steps++
			switch instruction {
			case 'L':
				node = d.graph[node].leftEdge
			case 'R':
				node = d.graph[node].rightEdge
			}
			if strings.HasSuffix(node, "Z") {
				solutionChan <- steps
				return
			}
		}
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func (d *day8) GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func (d *day8) LCM(a, b int, integers ...int) int {
	result := a * b / d.GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = d.LCM(result, integers[i])
	}

	return result
}

type node struct {
	id        string
	leftEdge  string
	rightEdge string
}

func day8Part1() string {
	d := day8{
		input: "input/day8_input.txt",
		graph: map[string]node{},
	}
	d.parseInput()
	return fmt.Sprintf("%d", d.traverseGraph())
}

func day8Part2() string {
	d := day8{
		input: "input/day8_input.txt",
		graph: map[string]node{},
	}
	d.parseInput()
	return fmt.Sprintf("%d", d.traverseGraphLikeAGhost())
}
