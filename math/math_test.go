package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMod(t *testing.T) {
	cases := []struct {
		a, b     uint16
		expected uint16
	}{
		{0, 15, 0},
		{10, 5, 0},
		{3, 2, 1},
	}
	for _, tc := range cases {
		actual := Mod(tc.a, tc.b)
		assert.Equal(t, actual, tc.expected)
	}
}
