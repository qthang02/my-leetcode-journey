package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestFindPair(t *testing.T) {
	tc := []struct {
		input  []int
		k      int
		output int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
	}

	for _, c := range tc {
		name := fmt.Sprintf("input: %v, k: %d", c.input, c.k)

		t.Run(name, func(t *testing.T) {
			got := findPairs(c.input, c.k)

			if got != c.output {
				t.Errorf("got: %d, want: %d", got, c.output)
			}
		})
	}
}