package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Result struct {
	Line   string
	First  int
	Second int
}

func main() {
	inputFile, _ := inputFlags()

	puzzleInput, err := readFile(inputFile)
	if err != nil {
		panic(err)
	}

	var1, var2 := splitInput(puzzleInput)
	result := calcfreshIngredients(var1, var2)

	fmt.Println("Part 1:", result)
	fmt.Println("Part 2:", countFreshIDs(var1))

}

func splitInput(input []string) ([][2]int, []int) {
	emptyLine := ""
	newStart := 0
	inputRanges := [][2]int{}
	inputIDs := []int{}
	//inputMap := make(map[int]int)
	// Split input to grab fresh ingredient ranges and stop when hitting empty line
	for i := range input {
		if input[i] != emptyLine {
			//inputRanges = append(inputRanges, input[i])
			leftVal := strings.Split(input[i], "-")[0]
			rightVal := strings.Split(input[i], "-")[1]
			leftInt, _ := strconv.Atoi(leftVal)
			rightInt, _ := strconv.Atoi(rightVal)
			inputRanges = append(inputRanges, [2]int{leftInt, rightInt})
		} else {
			newStart = i + 1
			break
		}
	}
	// Grab remaining lines as ingredient IDs
	for j := newStart; j < len(input); j++ {
		convInt, _ := strconv.Atoi(input[j])
		inputIDs = append(inputIDs, convInt)
	}

	return inputRanges, inputIDs
}

func calcfreshIngredients(ranges [][2]int, ids []int) int {
	resultCount := 0
	for _, id := range ids {
		isFresh := false
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				isFresh = true
				break
			}
		}
		if isFresh {
			resultCount++
		}
	}

	return resultCount
}

func countFreshIDs(ranges [][2]int) int {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start value (r[0])
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// Walk through and merge overlaps
	curStart := ranges[0][0]
	curEnd := ranges[0][1]
	total := 0

	for i := 1; i < len(ranges); i++ {
		start := ranges[i][0]
		end := ranges[i][1]

		if start <= curEnd+1 {
			// Overlaps or directly touches current range, extend it
			if end > curEnd {
				curEnd = end
			}
		} else {
			// No overlap, close off current range and start a new one
			total += curEnd - curStart + 1
			curStart = start
			curEnd = end
		}
	}

	// Add the last merged range
	total += curEnd - curStart + 1

	return total
}
