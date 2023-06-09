package Array_String_HashMap

func mergeAlternately(word1 string, word2 string) string {
	rs := ""

	if len(word1) == 0 {
		return word2
	}
	if len(word2) == 0 {
		return word1
	}

	if len(word1) == len(word2) {
		for i := 0; i < len(word1); i++ {
			rs += string(word1[i]) + string(word2[i])
		}
	} else if len(word1) > len(word2) {
		for i := 0; i < len(word1); i++ {
			if len(word2) <= i {
				rs += string(word1[i])
			} else {
				rs += string(word1[i]) + string(word2[i])
			}
		}
	} else {
		for i := 0; i < len(word2); i++ {
			if len(word1) <= i {
				rs += string(word2[i])
			} else {
				rs += string(word1[i]) + string(word2[i])
			}
		}
	}

	return rs
}
