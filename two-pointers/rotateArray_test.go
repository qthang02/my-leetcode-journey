package two_pointers

import (
	"fmt"
	"testing"
)

func TestRotate(m *testing.T) {
	tc := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, c := range tc {
		name := fmt.Sprintf("input: %v, k: %d", c.input, c.k)

		m.Run(name, func(m *testing.T) {
			got := rotate(c.input, c.k)

			if !checkValue(got, c.output) {
				m.Errorf("got: %v, want: %v", got, c.output)
			}
		})
	}
}

func checkValue(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
