package binary_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	tc := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, 2, -1},
		{[]int{5, 6}, 5, 0},
	}

	for _, c := range tc {
		name := fmt.Sprintf("%v, %v", c.nums, c.target)

		t.Run(name, func(t *testing.T) {
			got := search(c.nums, c.target)
			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}
