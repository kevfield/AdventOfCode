package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Grab user choices
	inputFile, _ := inputFlags()
	puzzleInput, _ := readfileasString(inputFile)

	// Pull in puzzle input
	rules, updates := parseInput(puzzleInput)
	total := 0
	incorrectTotal := 0

	// Part 1: Calculate sum of middle pages for valid updates.
	for _, update := range updates {
		if isValid(update, rules) {
			total += findMiddlePage(update)
		}
	}

	fmt.Println("Sum of middle pages (Part 1):", total)

	// Part 2: Process updates and calculate the sum of middle pages for reordered updates.
	for _, update := range updates {
		if !isValid(update, rules) {
			reordered := reorderUpdate(update, rules)
			incorrectTotal += findMiddlePage(reordered)
		}
	}

	fmt.Println("Sum of middle pages (Part 2):", incorrectTotal)
}

func parseInput(input string) (map[int][]int, [][]int) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	ruleLines := strings.Split(sections[0], "\n")
	updateLines := strings.Split(sections[1], "\n")

	// Parse rules into a map
	rules := make(map[int][]int)
	for _, line := range ruleLines {
		parts := strings.Split(line, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		rules[a] = append(rules[a], b)
	}
	// Parse updates into a slice of slices
	var updates [][]int
	for _, line := range updateLines {
		parts := strings.Split(line, ",")
		var update []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func isValid(update []int, rules map[int][]int) bool {
	// Create a map to store the index of each page in the update
	indexMap := make(map[int]int)
	for i, page := range update {
		indexMap[page] = i
	}

	// Validate each rule
	for a, dependencies := range rules {
		for _, b := range dependencies {
			if idxA, okA := indexMap[a]; okA {
				if idxB, okB := indexMap[b]; okB {
					if idxA > idxB {
						return false
					}
				}
			}
		}
	}

	return true
}

func findMiddlePage(update []int) int {
	mid := len(update) / 2
	return update[mid]
}

func reorderUpdate(update []int, rules map[int][]int) []int {
	// Create in-degree map and adjacency list.
	inDegree := make(map[int]int)
	adjList := make(map[int][]int)

	// Build graph from rules.
	for a, dependencies := range rules {
		for _, b := range dependencies {
			inDegree[b]++                      // b has an incoming edge from a
			adjList[a] = append(adjList[a], b) // a points to b
		}
	}

	// Initialize queue with nodes that have no dependencies (in-degree 0).
	queue := []int{}
	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	// Perform topological sort: process nodes with no dependencies first
	sorted := []int{}
	visited := make(map[int]bool)

	// Perform the topological sort process
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)
		visited[node] = true

		// Process each dependent (neighbor) of the node
		for _, neighbor := range adjList[node] {
			inDegree[neighbor]-- // Decrease the in-degree of neighbor
			if inDegree[neighbor] == 0 && !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	// Now, add any pages that were not processed by the topological sort
	// These pages don't have any dependencies or were already placed.
	for _, page := range update {
		if !visited[page] {
			sorted = append(sorted, page)
		}
	}

	return sorted
}
