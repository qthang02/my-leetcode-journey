package binary_search

import (
	"fmt"
	"testing"
)

func TestSearchRotated(t *testing.T) {
	tc := []struct {
		nums   []int
		target int
		want   int
	} {
		{[]int{4,5,6,7,0,1,2}, 0, 4},
		{[]int{4,5,6,7,0,1,2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for _, c := range tc {
		name := fmt.Sprintf("%v, %v", c.nums, c.target)

		t.Run(name, func(t *testing.T) {
			got := searchRotated(c.nums, c.target)
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}