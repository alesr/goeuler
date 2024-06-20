package euler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		given    int
		expected int
	}{
		{
			name:     "example",
			given:    6,
			expected: 13,
		},
		{
			name:     "solution",
			given:    10001,
			expected: 104743,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			observed := nthPrime(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}

func TestIsPrime(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    int
		expected bool
	}{
		{
			given:    2,
			expected: true,
		},
		{
			given:    3,
			expected: true,
		},
		{
			given:    4,
			expected: false,
		},
		{
			given:    5,
			expected: true,
		},
		{
			given:    6,
			expected: false,
		},
		{
			given:    7,
			expected: true,
		},
		{
			given:    8,
			expected: false,
		},
		{
			given:    9,
			expected: false,
		},
		{
			given:    10,
			expected: false,
		},
		{
			given:    11,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: '%d'", tc.given), func(t *testing.T) {
			t.Parallel()

			observed := isPrime(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
