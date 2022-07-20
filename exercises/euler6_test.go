package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSquareDiff(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    int
		expected int
	}{
		{
			given:    10,
			expected: 2640,
		},
		{
			given:    100,
			expected: 25164150,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: '%d', expected: '%d'", tc.given, tc.expected), func(t *testing.T) {
			observed := sumSquareDiff(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
