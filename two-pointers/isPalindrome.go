package two_pointers

import "strings"

func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	rs := ""
	c := ""

	for i := len(str) - 1; i >= 0; i-- {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= '0' && str[i] <= '9') {
			rs += string(str[i])
		}
	}

	for i := 0; i < len(s); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= '0' && str[i] <= '9') {
			c += string(str[i])
		}
	}

	return rs == c
}
