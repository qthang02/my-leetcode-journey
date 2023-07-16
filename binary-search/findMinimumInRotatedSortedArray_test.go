package binary_search

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	tc := []struct {
		input []int
		want  int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 1}, 1},
	}

	for _, c := range tc {
		name := fmt.Sprintf("findMin(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			got := findMin(c.input)

			if got != c.want {
				t.Fatalf("findMin(%v) = %v, want %v", c.input, got, c.want)
			}
		})
	}
}