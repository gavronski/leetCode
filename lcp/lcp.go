package lcp

// find longest common prefix in array of strings
func LongestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}

	// iterate through first string charcters
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		// compare n character of first string with n character of rest strings
		for j := 1; j < len(strs); j++ {
			// return prefix when characters are not equal
			if strs[j][i] != char || len(strs[j]) < len(strs[0]) {
				return prefix
			}
		}
		prefix = prefix + string(char)
	}

	return prefix
}
