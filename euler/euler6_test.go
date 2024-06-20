package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSquareDiff(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		given    int
		expected int
	}{
		{
			name:     "example",
			given:    10,
			expected: 2640,
		},
		{
			name:     "solution",
			given:    100,
			expected: 25164150,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			observed := sumSquareDiff(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
