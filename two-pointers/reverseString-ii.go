package two_pointers

func reverseStr(s string, k int) string {
	tmp := []byte(s)

	for i := 0; i < len(s); i += 2 * k {
		r := i + k - 1
		if len(s)-1 < r {
			r = len(s) - 1
		}

		for l := i; l < r; l, r = l+1, r-1 {
			tmp[l], tmp[r] = tmp[r], tmp[l]
		}
	}

	return string(tmp)
}