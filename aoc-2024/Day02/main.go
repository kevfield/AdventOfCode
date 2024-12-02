package main

import "fmt"

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := ReadFileAsNestedIntSlice(inputFile)

	fmt.Println("Part 1", ascorDesc(puzzleInput))
}

func validTests(inttoTest []int) bool {
	// must increment less than 4
	// must be continually ascending or descending
	// must not have two numbers in a row that are the same
	var ascCount, descCount, validCount int
	for i := 0; i < len(inttoTest)-1; i++ {
		if inttoTest[i] < inttoTest[i+1] {
			ascCount++
		} else {
			descCount++
		}
	}

	for j := 0; j < len(inttoTest)-1; j++ {
		if abs(inttoTest[j]-inttoTest[j+1]) < 4 && inttoTest[j] != inttoTest[j+1] {

			validCount++
		} else {
			validCount = 0
			break
		}
	}

	if validCount != 0 {
		if ascCount == 0 || descCount == 0 {
			return true
		}
	}
	return false
}

func ascorDesc(puzzleInput [][]int) int {
	resultCount := 0
	for _, v := range puzzleInput {
		if validTests(v) == true {
			fmt.Println(v)
			resultCount++
		}
	}
	return resultCount
}
