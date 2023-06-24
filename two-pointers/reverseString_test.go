package two_pointers

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	tc := []struct {
		input  []byte
		output []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, c := range tc {
		name := fmt.Sprintf("reverseString(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			got := reverseString(c.input)
			if !equalString(got, c.output) {
				t.Errorf("got %v, want %v", got, c.output)
			}
		})
	}
}

func equalString(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}