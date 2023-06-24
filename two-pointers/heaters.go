package two_pointers

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	distance := make([]int, len(houses))
	result := 0

	for i := range distance {
		distance[i] = math.MaxInt32
	}

	for i, j := 0, 0; i < len(houses) && j < len(heaters); {
		if houses[i] <= heaters[j] {
			distance[i] = heaters[j] - houses[i]
			i++
		} else {
			j++
		}
	}

	for i, j := len(houses) - 1, len(heaters) - 1; i >= 0 && j >= 0; {
		if houses[i] >= heaters[j] {
			distance[i] = min(distance[i], houses[i] - heaters[j])
			i--
		} else {
			j--
		}
	}

	for _, v := range distance {
		if v > result {
			result = v
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} 
	return b
}