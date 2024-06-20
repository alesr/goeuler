package euler

// Problem 3:

// Largest prime factor.
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?
// https://projecteuler.net/problem=3
func largestPrimeFactors(n int) int {
	var largest int
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			largest = i
			n = n / i
			i = 2
		}
	}
	return largest
}
