package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputFile, _ := inputFlags()
	puzzleInput, _ := readFile(inputFile)

	part1 := countZeroLandings(puzzleInput)
	part2 := countZeroCrossings(puzzleInput)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func countZeroLandings(lines []string) int {
	size := 100
	position := 50
	zeroCount := 0

	for _, line := range lines {
		if len(line) > 0 {
			direction := string(line[0])
			steps, _ := strconv.Atoi(line[1:])

			if direction == "L" {
				position = (position - steps%size + size) % size
			} else {
				position = (position + steps) % size
			}

			if position == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}

func countZeroCrossings(lines []string) int {
	size := 100
	position := 50
	zeroCount := 0

	for _, line := range lines {
		if len(line) > 0 {
			direction := string(line[0])
			steps, _ := strconv.Atoi(line[1:])

			if direction == "L" {
				for i := 0; i < steps; i++ {
					position--
					if position < 0 {
						position = size - 1
					}
					if position == 0 {
						zeroCount++
					}
				}
			} else {
				for i := 0; i < steps; i++ {
					position++
					if position >= size {
						position = 0
					}
					if position == 0 {
						zeroCount++
					}
				}
			}
		}
	}

	return zeroCount
}
