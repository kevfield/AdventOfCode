package main

import (
	"fmt"
)

func main() {
	// Grab user choices
	inputFile, _ := inputFlags()

	// Pull in puzzle input
	puzzleInput, _ := readFile(inputFile)
	fmt.Println(puzzleInput)

	fmt.Println("Part 1: ")

	fmt.Println("Part 2: ")

}
