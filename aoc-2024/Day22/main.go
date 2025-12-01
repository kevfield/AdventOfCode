package main

import (
	"fmt"
)

// Function to simulate the secret number evolution process
func nextSecretNumber(secret int) int {
	// Step 1: Multiply by 64, mix with XOR, and prune
	secret = (secret * 64) ^ secret
	secret = secret % 16777216

	// Step 2: Divide by 32, mix with XOR, and prune
	secret = (secret ^ (secret / 32)) % 16777216

	// Step 3: Multiply by 2048, mix with XOR, and prune
	secret = (secret * 2048) ^ secret
	secret = secret % 16777216

	return secret
}

func main() {
	// Initial secret numbers for each buyer
	inputFile, _ := inputFlags()
	puzzleInput, _ := readFileAsIntSlice(inputFile)

	totalSum := 0

	// For each buyer, simulate the secret number generation
	for _, initialSecret := range puzzleInput {
		secret := initialSecret
		// Simulate 2000 secret numbers for the buyer
		for i := 0; i < 2000; i++ {
			secret = nextSecretNumber(secret)
		}
		// Add the 2000th secret number to the total sum
		totalSum += secret
	}

	// Output the result
	fmt.Println("Total sum of the 2000th secret numbers:", totalSum)
}
