package two_pointers

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tc := []struct {
		s string
		t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "ahbgdc", true},
	}

	for _, tt := range tc {
		name := fmt.Sprintf("%v, %v", tt.s, tt.t)
		
		t.Run(name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}