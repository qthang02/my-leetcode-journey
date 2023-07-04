package two_pointers

func maxArea(height []int) int {
	rs, area, left, right := 0, 0, 0, len(height)-1
	for left < right {
		area = min(height[left], height[right]) * (right - left)
		rs = max(rs, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}