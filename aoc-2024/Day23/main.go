package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	networkMap := []string{
		"kh-tc", "qp-kh", "de-cg", "ka-co", "yn-aq", "qp-ub",
		"cg-tb", "vc-aq", "tb-ka", "wh-tc", "yn-cg", "kh-ub",
		"ta-co", "de-ta", "tc-td", "tb-wq", "wh-td", "ta-ka",
		"td-qp", "aq-cg", "wq-ub", "ub-vc", "de-ta", "wq-aq",
		"wq-vc", "wh-yn", "ka-de", "kh-ta", "co-tc", "wh-qp",
		"tb-vc", "td-yn",
	}

	// Create an adjacency list
	adjList := make(map[string]map[string]bool)
	for _, connection := range networkMap {
		parts := strings.Split(connection, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		if adjList[a] == nil {
			adjList[a] = make(map[string]bool)
		}
		if adjList[b] == nil {
			adjList[b] = make(map[string]bool)
		}
		adjList[a][b] = true
		adjList[b][a] = true
	}

	// Find all unique triangles
	triangleSet := map[string]bool{}
	for a, neighbors := range adjList {
		for b := range neighbors {
			if b <= a {
				continue // Avoid duplicates
			}
			for c := range adjList[b] {
				if c <= b || !adjList[a][c] {
					continue
				}
				// Sort the triangle members to ensure uniqueness
				triangle := []string{a, b, c}
				sort.Strings(triangle)
				triangleKey := strings.Join(triangle, ",")
				triangleSet[triangleKey] = true
			}
		}
	}

	// Filter triangles with at least one name starting with 't'
	validTriangles := 0
	for triangleKey := range triangleSet {
		names := strings.Split(triangleKey, ",")
		for _, name := range names {
			if strings.HasPrefix(name, "t") {
				validTriangles++
				break
			}
		}
	}

	fmt.Printf("Total unique triangles: %d\n", len(triangleSet))
	fmt.Printf("Triangles with at least one 't'-prefix: %d\n", validTriangles)
}
