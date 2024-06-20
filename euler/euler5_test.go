package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestMultiple(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		given    int
		expected int
	}{
		{
			name:     "example",
			given:    10,
			expected: 2520,
		},
		{
			name:     "solution",
			given:    20,
			expected: 232792560,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			observed := smallestMultiple(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
