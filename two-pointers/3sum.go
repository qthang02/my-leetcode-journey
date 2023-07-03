package two_pointers

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)

	// Sort the given array
	// [-1, 0, 1, 2, -1, -4] ==> [ [-1, -1, 2], [-1, 0, 1] ],
	sort.Ints(nums)
	// [-4, -1, -1, 0, 1, 2]


	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1 // 1st element after num1Idx
		num3Idx := n - 1 	 // Last element
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}
