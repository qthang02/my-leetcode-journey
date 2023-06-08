package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tc := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{0, 4, 5, 0, 3, 6}, true},
	}

	for _, c := range tc {
		name := fmt.Sprintf("tc case: nums: %v", c.input)

		t.Run(name, func(t *testing.T) {
			rs := containsDuplicate(c.input)
			if rs != c.expected {
				t.Fatalf("expected: %v, got: %v", c.expected, rs)
			}
		})
	}
}
