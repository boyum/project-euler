package main

import "fmt"

func main() {
	// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	// Find the sum of all the multiples of 3 or 5 below 1000.

	const n int = 1000

	sum := 0

	for i := 0; i < n; i++ {
		if isMultipleOf3Or5(i) {
			sum += i
		}
	}

	fmt.Printf("The sum of all multiples of 3 or 5 below 1000 is: %d", sum)
}

func isMultipleOf3Or5(num int) bool {
	return isMultipleOf3(num) || isMultipleOf5(num)
}

func isMultipleOf3(num int) bool {
	return num%3 == 0
}

func isMultipleOf5(num int) bool {
	return num%5 == 0
}
