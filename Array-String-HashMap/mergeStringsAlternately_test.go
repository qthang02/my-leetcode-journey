package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tc := []struct {
		word1 string
		word2 string
		want  string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, c := range tc {
		name := fmt.Sprintf("mergeAlternately(%v, %v)", c.word1, c.word2)

		t.Run(name, func(t *testing.T) {
			got := mergeAlternately(c.word1, c.word2)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}
