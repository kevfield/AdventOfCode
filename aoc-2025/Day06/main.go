package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	First    int
	Second   int
	Third    int
	Operator string
}

func main() {
	inputFile, _ := inputFlags()
	//datastuff := Result{}
	puzzleInput, _ := readFile(inputFile)
	//groupMax := 4
	totalSum := 0

	row0 := strings.Fields(puzzleInput[0])
	row1 := strings.Fields(puzzleInput[1])
	row2 := strings.Fields(puzzleInput[2])
	row3 := strings.Fields(puzzleInput[3])
	row4 := strings.Fields(puzzleInput[4])

	for i := 0; i < len(row0); i++ {

		// convert string → int
		a, _ := strconv.Atoi(row0[i])
		b, _ := strconv.Atoi(row1[i])
		c, _ := strconv.Atoi(row2[i])
		d, _ := strconv.Atoi(row3[i])

		op := row4[i] // "+" or "*"

		// apply the operator twice: (a OP b) OP c
		temp1 := applyOp(a, b, op)
		temp2 := applyOp(temp1, c, op)
		vertSum := applyOp(temp2, d, op)

		totalSum += vertSum
		vertSum = 0
	}
	fmt.Println("Part 1:", totalSum)

}

func applyOp(a, b int, op string) int {
	if op == "+" {
		return a + b
	}
	if op == "*" {
		return a * b
	}
	panic("unknown operator: " + op)
}
