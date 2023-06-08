package Array_String_HashMap

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		_, ok := m[v]
		m[v] = false
		if ok {
			return true
		}
	}

	return false
}
