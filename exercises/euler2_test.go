package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenFibonacci(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    int
		expected int
	}{
		{
			given:    10,
			expected: 10,
		},
		{
			given:    4000000,
			expected: 4613732,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: %d, expected: %d", tc.given, tc.expected), func(t *testing.T) {
			observed := evenFibonacci(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
