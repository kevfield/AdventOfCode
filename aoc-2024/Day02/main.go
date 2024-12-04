package main

import "fmt"

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()
	// pull in puzzle input
	puzzleInput, _ := ReadFileAsNestedIntSlice(inputFile)

	fmt.Println("Part 1: ", ascorDesc(puzzleInput))
}

func validTests(inttoTest []int) (bool, int) {
	// must increment less than 4
	// must be continually ascending or descending
	// must not have two numbers in a row that are the same
	var ascCount, descCount, validMatch, validAbs, validascDesc, failures int
	for i := 0; i < len(inttoTest)-1; i++ {
		if inttoTest[i] < inttoTest[i+1] {
			ascCount++
		} else {
			descCount++
		}
	}

	for j := 0; j < len(inttoTest)-1; j++ {
		if abs(inttoTest[j]-inttoTest[j+1]) < 4 {
			validAbs++
		} else {
			validAbs = 0
			failures++
			break
		}
	}
	for j := 0; j < len(inttoTest)-1; j++ {
		if inttoTest[j] != inttoTest[j+1] {
			validMatch++
		} else {
			validMatch = 0
			failures++
			break
		}
	}
	if ascCount == 0 || descCount == 0 {
		validascDesc++
	} else {
		validascDesc = 0
		failures++
	}

	if validAbs != 0 && validMatch != 0 && validascDesc != 0 {
		return true, failures
	}
	return false, failures
}

func createSlice(importedSlice []int, index int) []int {
	if index < 0 || index >= len(importedSlice) {
		// Return the original slice if the index is out of bounds
		return importedSlice
	}

	// Create a new slice excluding the element at the specified index
	newSlice := append([]int{}, importedSlice[:index]...)   // Copy elements before index
	newSlice = append(newSlice, importedSlice[index+1:]...) // Copy elements after index
	return newSlice
}

func ascorDesc(puzzleInput [][]int) int {
	resultCount := 0
	p2resultCount := 0
	//var newslice []int
	for _, v := range puzzleInput {
		valid1, _ := validTests(v)
		if valid1 == true {
			resultCount++
		} else {
			for i := 0; i < len(v); i++ {
				newSlice := createSlice(v, i)
				if valid2, fails := validTests(newSlice); valid2 == true && fails != 1 {
					p2resultCount++
					break
				}
			}
		}
	}

	fmt.Println("P2 Result: ", p2resultCount+resultCount)
	return resultCount
}
