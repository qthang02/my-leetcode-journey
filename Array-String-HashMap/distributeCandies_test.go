package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{6, 6, 6, 6}, 1},
	}

	for _, c := range tc {
		name := fmt.Sprintf("distributeCandies(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			got := distributeCandies(c.input)
			if got != c.output {
				t.Errorf("got %d, want %d", got, c.output)
			}
		})
	}
}