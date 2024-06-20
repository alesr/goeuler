package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPrimeFactors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		given    int
		expected int
	}{
		{
			name:     "example",
			given:    13195,
			expected: 29,
		},
		{
			name:     "solution",
			given:    600851475143,
			expected: 6857,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			observed := largestPrimeFactors(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
