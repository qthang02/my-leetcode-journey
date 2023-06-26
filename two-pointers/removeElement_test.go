package two_pointers

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tc := []struct {
		input    []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, c := range tc {
		name := fmt.Sprintf("removeElement(%v, %d)", c.input, c.val)

		t.Run(name, func(t *testing.T) {
			got := removeElement(c.input, c.val)

			if got != c.expected {
				t.Errorf("got %d, expected %d", got, c.expected)
			}
		})
	}
}