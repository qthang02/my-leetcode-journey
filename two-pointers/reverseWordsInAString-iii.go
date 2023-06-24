package two_pointers

import "strings"

func reverseWords(s string) string {
	var result []string

	for _, v := range strings.Split(s, " ") {
		tmp := []byte(v)

		for i, j := 0, len(tmp) - 1; i < j; {
			tmp[i], tmp[j] = tmp[j], tmp[i]
			i++
			j--
		}

		result = append(result, string(tmp))
	}

	return strings.Join(result, " ")
}