package two_pointers

import (
	"fmt"
	"testing"
)

func TestFindRadius(t *testing.T) {
	tc :=  []struct {
		houses []int
		heaters []int
		output int
	}{
		{[]int{1,2,3}, []int{2}, 1},
		{[]int{1,2,3,4}, []int{1,4}, 1},
		{[]int{1,5}, []int{2}, 3},
	}

	for _, c := range tc {
		name := fmt.Sprintf("findRadius(%v, %v)", c.houses, c.heaters)

		t.Run(name, func(t *testing.T) {
			got := findRadius(c.houses, c.heaters)
			if got != c.output {
				t.Errorf("got %d, want %d", got, c.output)
			}
		})
	}
}