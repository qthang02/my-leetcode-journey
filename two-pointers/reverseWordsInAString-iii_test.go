package two_pointers

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tc := []struct {
		input  string
		output string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
	}

	for _, c := range tc {
		name := fmt.Sprintf("reverseWords(%s)", c.input)

		t.Run(name, func(t *testing.T) {
			got := reverseWords(c.input)
			if got != c.output {
				t.Errorf("got %s, want %s", got, c.output)
			}
		})
	}
}