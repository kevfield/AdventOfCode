package main

import (
	"fmt"
	"regexp"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	matchMul(puzzleInput)

}

func matchMul(matchInput []string) {
	var firstInt, secondInt, mulTotal, mulSingle int
	regMul := regexp.MustCompile(`mul\(-?\d+,-?\d+\)`)
	extractedMul := regMul.FindAllString(matchInput[0], -1)

	for i := 0; i < len(extractedMul); i++ {
		fmt.Sscanf(extractedMul[i], "mul(%d,%d)", &firstInt, &secondInt)
		mulSingle = firstInt * secondInt
		mulTotal += mulSingle
	}
	fmt.Println("Part 1: ", mulTotal)

}
