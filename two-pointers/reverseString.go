package two_pointers

func reverseString(s []byte) []byte {
	if len(s) == 0 {
		return s
	}
	i, j := 0, len(s)-1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}