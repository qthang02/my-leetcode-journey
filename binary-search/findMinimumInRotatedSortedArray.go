package binary_search

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	min := nums[0]
	for l <= r {
		m := (l + r) / 2
		if nums[m] >= min {
			l = m + 1
		} else {
			if nums[m] < min {
				min = nums[m]
				r = m - 1
			}
		}
	}
	return min
}