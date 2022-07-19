package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplesOf3Or5(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		given    int
		expected int
	}{
		{
			given:    10,
			expected: 23,
		},
		{
			given:    1000,
			expected: 233168,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given: %d, expected: %d", tc.given, tc.expected), func(t *testing.T) {
			observed := multiplesOf3Or5(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
