package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tc := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"cider", "cried", true},
		{"aacc", "ccac", false},
	}

	for _, c := range tc {
		name := fmt.Sprintf("isAnagram(%q, %q)", c.s, c.t)

		t.Run(name, func(t *testing.T) {
			got := isAnagram(c.s, c.t)
			if got != c.expected {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}
