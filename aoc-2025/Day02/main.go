package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputFile, _ := inputFlags()

	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	part1, part2, err := sumInvalidIDs(string(content), ",")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func findDoubledDigits(start, end int) []int {
	var results []int

	for num := start; num <= end; num++ {
		s := strconv.Itoa(num)

		if len(s)%2 != 0 {
			continue
		}

		half := len(s) / 2
		if s[:half] == s[half:] {
			results = append(results, num)
		}
	}

	return results
}

func findRepeatedDigits(start, end int) []int {
	var results []int

	for num := start; num <= end; num++ {
		s := strconv.Itoa(num)

		for length := 1; length <= len(s)/2; length++ {
			if len(s)%length != 0 {
				continue
			}

			pattern := s[:length]
			repeats := len(s) / length

			if strings.Repeat(pattern, repeats) == s {
				results = append(results, num)
				break
			}
		}
	}

	return results
}

func sumInvalidIDs(input string, delimiter string) (int, int, error) {
	ranges, err := parseRanges(input)
	if err != nil {
		return 0, 0, err
	}

	part1 := 0
	part2 := 0

	for _, r := range ranges {
		for _, num := range findDoubledDigits(r[0], r[1]) {
			part1 += num
		}
		for _, num := range findRepeatedDigits(r[0], r[1]) {
			part2 += num
		}
	}

	return part1, part2, nil
}
