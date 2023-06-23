package two_pointers

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tc := []struct {
		input []int
		want  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, c := range tc {
		name := fmt.Sprintf("intput: %v", c.input)

		t.Run(name, func(t *testing.T) {
			got := moveZeroes(c.input)

			if !equal(got, c.want) {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}

func equal(a, b []int) bool {
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
