package main

import (
	"fmt"
	"math/big"
)

type Result struct {
	Line   string
	First  int
	Second int
}

// Part 1: Find two digits that form the largest two-digit number
func findTwoHighest(lines []string) ([]Result, error) {
	results := make([]Result, 0, len(lines))

	for _, s := range lines {
		if len(s) < 2 {
			continue
		}

		maxJoltage := -1
		first, second := 0, 0

		for i := 0; i < len(s)-1; i++ {
			if s[i] < '0' || s[i] > '9' {
				return nil, fmt.Errorf("non-digit character found in line: %s", s)
			}
			d1 := int(s[i] - '0')

			for j := i + 1; j < len(s); j++ {
				if s[j] < '0' || s[j] > '9' {
					return nil, fmt.Errorf("non-digit character found in line: %s", s)
				}
				d2 := int(s[j] - '0')

				joltage := d1*10 + d2
				if joltage > maxJoltage {
					maxJoltage = joltage
					first = d1
					second = d2
				}
			}
		}

		results = append(results, Result{
			Line:   s,
			First:  first,
			Second: second,
		})
	}

	return results, nil
}

// Part 2: Find N digits that form the largest N-digit number
func findMaxJoltage(s string, numDigits int) string {
	if len(s) <= numDigits {
		return s
	}

	result := make([]byte, 0, numDigits)
	startPos := 0

	for i := 0; i < numDigits; i++ {
		digitsRemaining := numDigits - i - 1
		lastValidPos := len(s) - digitsRemaining - 1

		bestPos := startPos
		for j := startPos; j <= lastValidPos; j++ {
			if s[j] > s[bestPos] {
				bestPos = j
			}
		}

		result = append(result, s[bestPos])
		startPos = bestPos + 1
	}

	return string(result)
}

func findMaxJoltages(lines []string, numDigits int) []string {
	results := make([]string, 0, len(lines))

	for _, s := range lines {
		if len(s) == 0 {
			continue
		}
		results = append(results, findMaxJoltage(s, numDigits))
	}

	return results
}

func main() {
	inputFile, _ := inputFlags()

	puzzleInput, err := readFile(inputFile)
	if err != nil {
		panic(err)
	}

	// Part 1
	results, err := findTwoHighest(puzzleInput)
	if err != nil {
		panic(err)
	}

	part1 := 0
	for _, r := range results {
		part1 += r.First*10 + r.Second
	}
	fmt.Println("Part 1:", part1)

	// Part 2
	joltages := findMaxJoltages(puzzleInput, 12)

	part2 := big.NewInt(0)
	for _, j := range joltages {
		n := new(big.Int)
		n.SetString(j, 10)
		part2.Add(part2, n)
	}
	fmt.Println("Part 2:", part2.String())
}
