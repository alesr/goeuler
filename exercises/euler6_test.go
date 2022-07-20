package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSquareDiff(t *testing.T) {
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
		observed := sumSquareDiff(tc.given)
		assert.Equal(t, tc.expected, observed)
	}
}
