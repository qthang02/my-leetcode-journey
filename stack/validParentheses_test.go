package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tc := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"{[()]}", true},
		{"{[()]}{[()]}", true},
	}

	for _, c := range tc {
		name := fmt.Sprintf("%s", c.input)

		t.Run(name, func(t *testing.T) {
			got := isValid(c.input)
			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}