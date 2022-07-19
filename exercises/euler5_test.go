package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestMultiple(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    int
		expected int
	}{
		{
			given:    10,
			expected: 2520,
		},
		{
			given:    20,
			expected: 232792560,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: '%d', expected: '%d'", tc.given, tc.expected), func(t *testing.T) {
			observed := smallestMultiple(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
