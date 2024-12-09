package main

import (
	"fmt"
)

func main() {
	// Grab user choices
	inputFile, _ := inputFlags()

	// Pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	fmt.Println("Part 1: ", findXmas(puzzleInput))

	fmt.Println("Part 2: ", findXMasPattern(puzzleInput))

}

// findXmas checks if the string "XMAS" exists in any direction within the input
func findXmas(input []string) int {
	var xmasCount int
	rows := len(input)

	dCols := len(input[0])

	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Vertical and horizontal
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonal
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < dCols; y++ {
			if input[x][y] == 'X' {
				for _, dir := range directions {
					if checkXmas(input, x, y, dir[0], dir[1]) {
						xmasCount++
					}
				}
			}
		}
	}

	return xmasCount
}

// checkXmas verifies if "XMAS" can be formed starting from (x, y) in the given direction (dx, dy)
func checkXmas(input []string, x, y, dx, dy int) bool {
	target := "XMAS"
	rows := len(input)
	cols := len(input[0])

	for i := 0; i < len(target); i++ {
		nx, ny := x+(dx*i), y+(dy*i)
		if nx < 0 || ny < 0 || nx >= rows || ny >= cols || input[nx][ny] != target[i] {
			return false
		}
	}

	return true
}

func findXMasPattern(input []string) int {
	var xmaspatCount int
	rows := len(input)
	cols := len(input[0])

	// Iterate through the grid (no border exclusion)
	for x := 1; x < rows-1; x++ { // Start from row 1 to row-1 to avoid boundary checks
		for y := 1; y < cols-1; y++ { // Start from col 1 to col-1 to avoid boundary checks
			// Check if the current position is the center 'A'
			if input[x][y] == 'A' {
				// Check the four surrounding positions (diagonal neighbors)
				topLeft := input[x-1][y-1]
				topRight := input[x-1][y+1]
				bottomLeft := input[x+1][y-1]
				bottomRight := input[x+1][y+1]

				//// Debugging output to check the characters around the center 'A'
				//fmt.Printf("Checking center 'A' at (%d, %d)\n", x, y)
				//fmt.Printf("TopLeft: %c, TopRight: %c, BottomLeft: %c, BottomRight: %c\n", topLeft, topRight, bottomLeft, bottomRight)

				// Check if the neighbors form the "X-MAS" pattern in any orientation
				if (topLeft == 'M' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'S') ||
					(topLeft == 'S' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'M') ||
					(topLeft == 'M' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'S') ||
					(topLeft == 'S' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'M') {
					// Valid "X-MAS" pattern found
					//fmt.Printf("Pattern found at (%d, %d)\n", x, y)
					xmaspatCount++
				}
			}
		}
	}

	return xmaspatCount
}
