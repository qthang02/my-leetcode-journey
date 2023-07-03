package two_pointers

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tc := []struct {
		input  []int
		output [][]int
	} {
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, c := range tc {
		name := fmt.Sprintf("threeSum(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			got := threeSum(c.input)

			if !isSameSlice(got, c.output) {
				t.Errorf("got: %v, want: %v", got, c.output)
			}
		})
	}
}

func isSameSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !isSameArray(v, b[i]) {
			return false
		}
	}

	return true
}

func isSameArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}