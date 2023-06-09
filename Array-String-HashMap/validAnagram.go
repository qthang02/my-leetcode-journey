package Array_String_HashMap

import "strings"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		a := strings.Split(s, "")
		b := strings.Split(t, "")
		m := make(map[string]int)

		for _, v := range a {
			_, ok := m[v]
			if !ok {
				m[v] = 1
			} else {
				m[v]++
			}
		}

		for _, v := range b {
			_, ok := m[v]
			if !ok {
				return false
			} else {
				m[v]--
			}
		}

		for _, v := range m {
			if v > 0 {
				return false
			}
		}
	}

	return true
}
