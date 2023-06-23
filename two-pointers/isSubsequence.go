package two_pointers

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
			i++
			continue
		}
		j++
	}

	return i == len(s)
}