package main

import (
	"fmt"
)

func main() {

	// Grab user choices
	inputFile, _ := inputFlags()

	// Pull in puzzle input
	puzzleInput, _ := readFile(inputFile)

	// Convert the map to a 2D slice of characters for easier manipulation
	mapData := make([][]rune, len(puzzleInput))
	for i, row := range puzzleInput {
		mapData[i] = []rune(row)
	}

	// Direction vectors for Up, Right, Down, Left
	directions := []struct{ dx, dy int }{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}

	// Find the initial position and direction
	var startX, startY, startDir int
	for i, row := range mapData {
		for j, cell := range row {
			if cell == '^' {
				startX, startY = j, i
				startDir = 0        // Up
				mapData[i][j] = '.' // Mark as empty now that we know the start position
				break
			}
		}
	}

	// Set to track distinct positions visited
	visited := make(map[string]bool)

	// Convert coordinates to a string for uniqueness
	coordsToStr := func(x, y int) string {
		return fmt.Sprintf("%d,%d", x, y)
	}

	// Initial position
	x, y := startX, startY
	dir := startDir

	// Simulate the guard's movement
	for {
		// Mark the current position as visited
		visited[coordsToStr(x, y)] = true

		// Check the next position
		newX, newY := x+directions[dir].dx, y+directions[dir].dy

		// Check if the position is out of bounds
		if newX < 0 || newX >= len(mapData[0]) || newY < 0 || newY >= len(mapData) {
			// Exit the map
			break
		}

		// If there's an obstacle, turn right (clockwise)
		if mapData[newY][newX] == '#' {
			dir = (dir + 1) % 4
			// Recalculate the new position after turning
			newX, newY = x+directions[dir].dx, y+directions[dir].dy
		}

		// Move the guard to the new position
		x, y = newX, newY
	}

	// Count distinct positions visited
	fmt.Println(len(visited))
}
