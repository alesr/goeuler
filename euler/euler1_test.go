package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplesOf3Or5(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		given    int
		expected int
	}{
		{
			name:     "example",
			given:    10,
			expected: 23,
		},
		{
			name:     "solution",
			given:    1000,
			expected: 233168,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			observed := multiplesOf3Or5(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
