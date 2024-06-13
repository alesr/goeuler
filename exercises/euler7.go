package exercises

import "math"

// Problem 7

// 10001st prime
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
// https://projecteuler.net/problem=7
func nthPrime(n int) int {
	var count int
	// We start at 3, so we're one iteration ahead.
	for i := 3; ; i += 2 {
		if isPrime(i) {
			count++
			if count == n-1 { // minus one iteration
				return i
			}
		}
	}
}

func isPrime(n int) bool {
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
