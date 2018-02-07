package main

import (
	"fmt"
	"time"
)

func main() {
	// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	// Find the largest palindrome made from the product of two 3-digit numbers.

	// Thought process:
	// The largest possible answer must be between
	// 10000 (100 * 100) and 998001 (999 * 999)
	//
	// Solution 1:
	// 1. Start with 999999
	// 2. Decrement the middle numbers, number is now 998899 (subtract 1100)
	// 3. Check if the number can be divided by any number between 999 and 100,
	//    and does the answer have exactly three digits?
	// 4. If no number can be found, decrement middle numbers again.
	// When middle numbers become 0, add 9 to them and decrement the second number from start and end (i.e. subtract 10010)

	start := time.Now()

	palindrome := 0
	factor1 := 0
	factor2 := 0
	counter := 0
	n := 999999

	for palindrome == 0 && counter < 999 {
		counter++

		if n == 100001 {
			n -= 2
		} else if n >= 100000 && n < 1000000 {
			if counter%100 == 0 {
				// Subtract first and last
				n -= 11
			} else if counter%10 == 0 {
				// Subtract second and second last
				n -= 110
			} else {
				// Subtract middle numbers
				n -= 1100
			}
		} else if n >= 10000 && n < 100000 {
			if counter%100 == 0 {
				// Subtract first and last
				n -= 11
			} else if counter%10 == 0 {
				// Subtract second and second last
				n -= 110
			} else {
				// Subtract middle numbers
				n -= 100
			}
		}

		for i := 999; i >= 500; i-- {
			if n%i == 0 && (n/i) < 1000 && (n/i) >= 100 {
				palindrome = n
				factor1 = i
				factor2 = n / i
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("The largest palindromic number is %d (%d*%d).\nExecution time: %s", palindrome, factor1, factor2, elapsed)
}
