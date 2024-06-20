package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenFibonacci(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		given    uint
		expected uint
	}{
		{
			name:     "example",
			given:    10,
			expected: 10,
		},
		{
			name:     "solution",
			given:    4000000,
			expected: 4613732,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			observed := evenFibonacci(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
