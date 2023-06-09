package Array_String_HashMap

import "strings"

func isAnagram(s string, t string) bool {
	str := strings.Split(s+t, "")
	m := make(map[string]int)

	for i, v := range str {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			if i < len(s) {
				m[v]++
			} else {
				m[v]--
			}
		}
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}
