package two_pointers

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tc := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tt := range tc {
		name := fmt.Sprintf("%v, %v", tt.numbers, tt.target)

		t.Run(name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("got: %v, want: %v", got, tt.want)
				}
			}
		})
	}
}