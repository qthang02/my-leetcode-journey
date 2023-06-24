package two_pointers

import (
	"fmt"
	"testing"
)

func TestReverseSt(t *testing.T) {
	tc := []struct {
		input  string
		k      int
		output string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
		{"abcdefg", 8, "gfedcba"},
	}

	for _, c := range tc {
		name := fmt.Sprintf("reverseStr(%v, %v)", c.input, c.k)

		t.Run(name, func(t *testing.T) {
			got := reverseStr(c.input, c.k)
			if got != c.output {
				t.Errorf("got %v, want %v", got, c.output)
			}
		})
	}
}