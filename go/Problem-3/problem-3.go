package main

import (
	"fmt"
	"time"
)

func main() {
	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143?

	num := 600851475143
	currentDivisor := 1
	largestPossibleFactor := num

	start := time.Now()

	for i := 1; i < largestPossibleFactor; i++ {
		currentDivisor++

		for isFactor(&currentDivisor, &largestPossibleFactor) && largestPossibleFactor >= 1 {
			largestPossibleFactor /= currentDivisor
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("The largest factor is %d. Execution time: %s", currentDivisor, elapsed)
}

func isFactor(factor *int, number *int) bool {
	return *number%*factor == 0
}
