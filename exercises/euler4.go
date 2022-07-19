package exercises

import "strconv"

// Problem 4

// Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
// https://projecteuler.net/problem=4
func largestPalindromeProduct() int {
	var max int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j
			if isPalindrome(strconv.Itoa(product)) {
				if product > max {
					max = product
				}
			}
		}
	}
	return max
}

func isPalindrome(s string) bool {
	return s == reverseString(s)
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
