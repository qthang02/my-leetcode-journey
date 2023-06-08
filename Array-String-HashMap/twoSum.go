package Array_String_HashMap

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		key := target - nums[i]

		if store[key] > 0 {
			return []int{store[key] - 1, i}
		}

		store[nums[i]] = i + 1
	}

	return nil
}
