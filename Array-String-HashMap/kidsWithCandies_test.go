package Array_String_HashMap

import (
	"fmt"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {

	tc := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
		{[]int{2, 8, 7}, 1, []bool{false, true, true}},
	}

	for _, c := range tc {
		name := fmt.Sprintf("%v, %v", c.candies, c.extraCandies)

		t.Run(name, func(t *testing.T) {
			got := kidsWithCandies(c.candies, c.extraCandies)
			if len(got) != len(c.want) {
				t.Fatalf("got %v, want %v", got, c.want)
			}
			for i := range got {
				if got[i] != c.want[i] {
					t.Fatalf("got %v, want %v", got, c.want)
				}
			}
		})
	}
}
