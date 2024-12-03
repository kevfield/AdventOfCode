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

	matchMul(puzzleInput)       //part 1
	matchmuldoDont(puzzleInput) //part 2

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

func matchmuldoDont(dodontInput []string) {
	var firstInt, secondInt, mulTotal, mulSingle int
	mulProceed := true
	regmuldoDont := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	extractedmuldoDont := regmuldoDont.FindAllString(dodontInput[0], -1)
	for i := range len(extractedmuldoDont) {
		switch extractedmuldoDont[i] {
		case "do()":
			mulProceed = true
		case "don't()":
			mulProceed = false
		default:
			if mulProceed {
				fmt.Sscanf(extractedmuldoDont[i], "mul(%d,%d)", &firstInt, &secondInt)
				mulSingle = firstInt * secondInt
				mulTotal += mulSingle
			}
		}
	}
	fmt.Println("Part 2: ", mulTotal)
}
