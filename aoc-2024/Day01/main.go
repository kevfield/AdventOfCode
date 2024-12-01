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

	fmt.Println("Part 1 =", splitintoSlices(puzzleInput))

}

func splitintoSlices(splitInput []string) int {
	var leftSlice, rightSlice []int
	var distances, calc int

	for i := 0; i < len(splitInput); i++ {
		splitLine := strings.Fields(splitInput[i])
		leftConv, _ := strconv.Atoi(splitLine[0])
		rightConv, _ := strconv.Atoi(splitLine[1])
		leftSlice = append(leftSlice, leftConv)
		rightSlice = append(rightSlice, rightConv)
	}

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	for j := 0; j < len(leftSlice); j++ {
		if leftSlice[j] > rightSlice[j] {
			calc = leftSlice[j] - rightSlice[j]
		} else {
			calc = rightSlice[j] - leftSlice[j]
		}
		distances = distances + calc
	}

	return distances
}
