package Array_String_HashMap

import "fmt"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		hash := [26]byte{}
		for _, c := range s {
			hash[c-'a']++
		}
		key := fmt.Sprintf("%x", hash)
		_, ok := groups[key]
		if !ok {
			groups[key] = []string{}
		}
		groups[key] = append(groups[key], s)
	}
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}