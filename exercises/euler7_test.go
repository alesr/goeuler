package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNstPrime(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    int
		expected int
	}{
		{
			given:    6,
			expected: 13,
		},
		{
			given:    10001,
			expected: 104743,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: '%d', expected: '%d'", tc.given, tc.expected), func(t *testing.T) {
			observed := nStPrime(tc.given)
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
		t.Run(fmt.Sprintf("given: '%d', expected: '%t'", tc.given, tc.expected), func(t *testing.T) {
			observed := isPrime(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
