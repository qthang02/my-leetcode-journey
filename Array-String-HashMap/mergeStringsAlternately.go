package Array_String_HashMap

func mergeAlternately(word1 string, word2 string) string {
	rs := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		rs += word1[i:j+1] + word2[i:j+1]
		i++
		j++
	}
	rs += word1[j:] + word2[j:]

	return rs
}
