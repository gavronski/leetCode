package firstoccurenceofneedle

import "log"

func StrStr(haystack string, needle string) int {
	firstChar := needle[0]
	index := -1

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == firstChar {
			if i <= len(haystack) && i+len(needle) <= len(haystack) {
				substr := haystack[i : len(needle)+i]
				log.Println(i, substr)
				if substr == needle {
					index = i
					break
				}
			}

		}
	}

	return index
}
