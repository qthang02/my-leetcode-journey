package two_pointers

func findPairs(nums []int, k int) int {
	result := 0

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for i, v := range m {
		if k == 0 && v >= 2 {
			result++
		} else if k > 0 && m[i+k] > 0 {
			result++
		}
	}

	return result
}
