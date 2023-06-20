package sliding_window

// "pwwkew", 3
// "abcabcbb", 3
func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	sLen := len(chars)

	if sLen < 2 {
		return sLen
	}

	maxCount := 0
	charDict := make(map[rune]int)
	var i, j = 0, 0

	for j < sLen {
		if c, ok := charDict[chars[j]]; ok {
			charDict[chars[j]] = c + 1
		} else {
			charDict[chars[j]] = 1
		}

		dictSize := len(charDict)
		windowSide := j - i + 1

		if dictSize == windowSide {
			if maxCount < windowSide {
				maxCount = windowSide
			}

			j++
			continue
		}

		if dictSize < windowSide {
			v := charDict[chars[i]] - 1
			if v == 0 {
				delete(charDict, chars[i])
			} else {
				charDict[chars[i]] = v
			}

			i++
			j++
			continue
		}
	}

	return maxCount
}
