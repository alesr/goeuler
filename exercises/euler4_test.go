package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPalindromeProduct(t *testing.T) {
	expected := 906609
	observed := largestPalindromeProduct()
	assert.Equal(t, expected, observed)
}

func TestReverseString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    string
		expected string
	}{
		{
			given:    "abc",
			expected: "cba",
		},
		{
			given:    "abcd",
			expected: "dcba",
		},
		{
			given:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		observed := reverseString(tc.given)
		assert.Equal(t, tc.expected, observed)
	}
}

func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    string
		expected bool
	}{
		{
			given: "1", expected: true,
		},
		{
			given: "30", expected: false,
		},
		{
			given: "9009", expected: true,
		},
		{
			given: "9008", expected: false,
		},
	}

	for _, tc := range testCases {
		observed := isPalindrome(tc.given)
		assert.Equal(t, tc.expected, observed)
	}
}
