package two_pointers

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tc := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"0P", false},
	}

	for _, c := range tc {
		name := fmt.Sprintf("isPalindrome(%q)", c.input)

		t.Run(name, func(t *testing.T) {
			got := isPalindrome(c.input)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}
