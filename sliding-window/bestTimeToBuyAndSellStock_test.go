package sliding_window

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}

	for _, c := range tc {
		name := fmt.Sprintf("%v", c.input)

		t.Run(name, func(t *testing.T) {
			res := maxProfit(c.input)

			if res != c.output {
				t.Errorf("want %d, got %d", c.output, res)
			}
		})
	}
}
