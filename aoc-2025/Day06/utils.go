package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read a file from an input and return into a slice of strings
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//lines = append(lines, "")
	return lines, scanner.Err()
}

// function to import a single line of input containing ranges like "2-4,6-8" and return a slice of integer pairs
func parseRanges(input string) ([][2]int, error) {
	parts := strings.Split(strings.TrimSpace(input), ",")
	ranges := make([][2]int, 0, len(parts))

	for _, part := range parts {
		bounds := strings.Split(part, "-")
		if len(bounds) != 2 {
			return nil, fmt.Errorf("invalid range: %s", part)
		}

		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, [2]int{start, end})
	}

	return ranges, nil
}

// ReadFileAsIntSlice reads a file and returns its contents as a slice of integers.
func ReadFileAsIntSlice(filename string) ([]int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var intSlice []int

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line into fields
		fields := strings.Fields(scanner.Text())
		for _, field := range fields {
			// Convert each field to an integer
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to convert '%s' to int: %w", field, err)
			}
			intSlice = append(intSlice, num)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return intSlice, nil
}

// use flags to get user input
func inputFlags() (string, string) {

	// variables declaration
	var flagfile string
	var partid string

	// flags declaration using flag package
	flag.StringVar(&flagfile, "file", "input.txt", "filename of the input data eg: input.txt")
	flag.StringVar(&partid, "part", "a", "use either a or b")

	flag.Parse() // after declaring flags we need to call it
	return flagfile, partid
}

func abs(x int) int {
	if x < 0 {
		return -x // Flip the sign if the number is negative
	}
	return x
}
