package sliding_window

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tc := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range tc {
		name := fmt.Sprintf("%s", tt.input)

		t.Run(name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %d, actual %d", tt.expected, actual)
			}
		})
	}
}
