package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	//splitintoSlices(puzzleInput)

	p1Answer, p2Answer := (calculateDistances(puzzleInput))
	fmt.Println("Part 1 =", p1Answer, "Part 2 =", p2Answer)

}

func calculateDistances(splitInput []string) (int, int) {
	var leftSlice, rightSlice []int
	var p1Distances, p1Calc, p2Calc, duplicates int

	for i := 0; i < len(splitInput); i++ {
		splitLine := strings.Fields(splitInput[i])
		leftConv, _ := strconv.Atoi(splitLine[0])
		rightConv, _ := strconv.Atoi(splitLine[1])
		leftSlice = append(leftSlice, leftConv)
		rightSlice = append(rightSlice, rightConv)
	}

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	// Part 1
	for j := 0; j < len(leftSlice); j++ {
		if leftSlice[j] > rightSlice[j] {
			p1Calc = leftSlice[j] - rightSlice[j]
		} else {
			p1Calc = rightSlice[j] - leftSlice[j]
		}
		p1Distances = p1Distances + p1Calc
	}

	// Part 2
	for k := 0; k < len(leftSlice); k++ {
		for l := 0; l < len(rightSlice); l++ {
			if leftSlice[k] == rightSlice[l] {
				duplicates++
			}
		}

		p2Calc = p2Calc + leftSlice[k]*duplicates

		// reset duplicates
		duplicates = 0
	}

	return p1Distances, p2Calc
}
