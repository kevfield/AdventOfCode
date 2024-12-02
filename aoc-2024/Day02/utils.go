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
func readFileAsString(filename string) ([]string, error) {
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

// ReadFileAsNestedIntSlice reads a file and returns its contents as a slice of slices of integers.
func ReadFileAsNestedIntSlice(filename string) ([][]int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var nestedIntSlice [][]int

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into fields
		fields := strings.Fields(line)
		var intSlice []int
		for _, field := range fields {
			// Convert each field to an integer
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to convert '%s' to int: %w", field, err)
			}
			intSlice = append(intSlice, num)
		}

		// Add the line's slice of integers to the main slice
		nestedIntSlice = append(nestedIntSlice, intSlice)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return nestedIntSlice, nil
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
