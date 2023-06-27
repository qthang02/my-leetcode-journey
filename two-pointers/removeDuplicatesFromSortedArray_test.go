package two_pointers

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, c := range tc {
		name := fmt.Sprintf("removeDuplicates(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			got := removeDuplicates(c.input)

			if got != c.output {
				t.Errorf("got %d, want %d", got, c.output)
			}
		})
	}
}