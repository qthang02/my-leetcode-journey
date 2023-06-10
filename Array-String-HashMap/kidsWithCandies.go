package Array_String_HashMap

func kidsWithCandies(candies []int, extraCandies int) []bool {
	rs := make([]bool, len(candies))
	m := make(map[int]bool)
	max := candies[0]

	for _, v := range candies {
		if max < v {
			max = v
		}
	}

	for i, v := range candies {

		vM, ok := m[v]

		if ok {
			rs[i] = vM
		} else {
			if v+extraCandies >= max {
				rs[i] = true
			} else {
				rs[i] = false
			}
			m[v] = rs[i]
		}
	}
	return rs
}
