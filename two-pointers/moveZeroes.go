package two_pointers

func moveZeroes(nums []int) []int {
	n := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}

	return nums
}
