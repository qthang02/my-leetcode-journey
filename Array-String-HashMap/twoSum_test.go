package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tc := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}

	for _, c := range tc {
		name := fmt.Sprintf("test case: nums: %v, target: %v", c.nums, c.target)

		t.Run(name, func(t *testing.T) {
			rs := twoSum(c.nums, c.target)
			fmt.Println(rs)
		})
	}
}
