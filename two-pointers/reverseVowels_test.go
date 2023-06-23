package two_pointers

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	tc := []struct {
		input string
		want  string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	}

	for _, c := range tc {
		name := fmt.Sprintf("input: %v", c.input)

		t.Run(name, func(t *testing.T) {
			got := reverseVowels(c.input)

			if got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}