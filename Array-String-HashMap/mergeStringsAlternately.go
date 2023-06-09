package Array_String_HashMap

func mergeAlternately(word1 string, word2 string) string {
	rs := ""
	i := 0

	for i < len(word1) && i < len(word2) {
		rs += word1[i:i+1] + word2[i:i+1]
		i++
	}
	rs += word1[i:] + word2[i:]

	return rs
}
