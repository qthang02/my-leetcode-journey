package two_pointers

func reverseVowels(s string) string {
	tmp := []byte(s)

	i, j := 0, len(s)-1
	for i < j {
		if isVowel(tmp[i]) && isVowel(tmp[j]) {
			tmp[i], tmp[j] = tmp[j], tmp[i]
			i++
			j--
		}
		if !isVowel(tmp[i]) {
			i++
			continue
		}
		if !isVowel(tmp[j]) {
			j--
			continue
		}
	}

	return string(tmp)
}

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	} else if b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}

	return false
}