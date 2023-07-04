package two_pointers

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, c := range tc {
		name := fmt.Sprintf("maxArea(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			rs := maxArea(c.input)
			if rs != c.output {
				t.Errorf("got %d, want %d", rs, c.output)
			}
		})
	}
}