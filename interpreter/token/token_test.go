package token_test

import (
	"testing"

	"github.com/ibraimgm/bfi/interpreter/token"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		token    rune
		expected bool
	}{
		{'>', true},
		{'<', true},
		{'+', true},
		{'-', true},
		{',', true},
		{'.', true},
		{'[', true},
		{']', true},
		{' ', false},
		{'#', false},
		{'a', false},
		{'\n', false},
		{'\t', false},
	}

	for i, v := range cases {
		result := token.IsValid(v.token)

		if result != v.expected {
			t.Errorf("Case %v, received \"%v\", expected: \"%v\"", i, result, v.expected)
		}
	}
}
