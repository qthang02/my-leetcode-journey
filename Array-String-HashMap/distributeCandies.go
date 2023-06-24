package Array_String_HashMap

func distributeCandies(candyType []int) int {
	m := make(map[int]bool)

	for _, v := range candyType {
		m[v] = true
	}

	return min(len(m), len(candyType)/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}