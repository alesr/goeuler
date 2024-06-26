package euler

// Problem 5:

// Smallest multiple.
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
// https://projecteuler.net/problem=5
func smallestMultiple(n int) int {
	smallestMultiple := 1

	for i := 1; i <= n; i++ {
		if smallestMultiple%i == 0 {
			if i == n {
				return smallestMultiple
			}
		} else {
			smallestMultiple++
			i = 1
		}
	}
	return smallestMultiple
}
