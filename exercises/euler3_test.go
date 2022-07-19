package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPrimeFactors(t *testing.T) {
	testCases := []struct {
		given    int
		expected int
	}{
		{
			given:    13195,
			expected: 29,
		},
		{
			given:    600851475143,
			expected: 6857,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: %d, expected: %d", tc.given, tc.expected), func(t *testing.T) {
			observed := largestPrimeFactors(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
