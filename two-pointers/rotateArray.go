package two_pointers

func rotate(nums []int, k int) []int {
	k = k % len(nums)

	reverse(&nums, 0, len(nums)-1)
	reverse(&nums, 0, k-1)
	reverse(&nums, k, len(nums)-1)

	return nums
}

func reverse(s *[]int, start, end int) {
	for start < end {
		(*s)[start], (*s)[end] = (*s)[end], (*s)[start]
	}
}
