package two_pointers

import "strconv"

func compress(chars []byte) int {
	scan := 0
	write := 0
	length := len(chars)

	for scan < length {
		chars[write] = chars[scan]
		count := 0

		for scan < length && chars[write] == chars[scan] {
			count++
			scan++
		}

		if count >= 2 {
			for _, v := range []byte(strconv.Itoa(count)) {
				chars[write] = v
				write++
			}
		}
		write++
	}

	return write
}