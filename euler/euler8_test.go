package euler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestProductInSeries(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		givenDigits   string
		givenAdjacent int
		expected      int
	}{
		{
			name:          "example",
			givenDigits:   input,
			givenAdjacent: 4,
			expected:      5832,
		},
		{
			name:          "problem",
			givenDigits:   input,
			givenAdjacent: 13,
			expected:      23514624000,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s, %d", tc.givenDigits, tc.givenAdjacent), func(t *testing.T) {
			observed := largestProductInSeries(tc.givenDigits, tc.givenAdjacent)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
