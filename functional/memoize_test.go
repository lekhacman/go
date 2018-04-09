package functional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoize(t *testing.T) {
	memoizedPlusTwo, _ := Memoize(func(i int) int {
		return i + 2
	})

	cases := []struct {
		input    int
		expected int
	}{
		{0, 2},
		{1, 3},
		{-6, -4},
		{16, 18},
		{1, 3},
	}

	for _, c := range cases {
		result := memoizedPlusTwo(c.input)
		assert.Equal(t, c.expected, result)
	}
}
