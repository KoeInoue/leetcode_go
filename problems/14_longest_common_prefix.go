package problem

import "strings"

func longestCommonPrefix(strs []string) string {
	var commonStr string
	str0Arr := strings.Split(strs[0], "")
	// each letter of first word
	for i := 0; i < len(str0Arr); i++ {
		var isMatchAll bool
		isMatchAll = true
		// each letters of other words
		for j := 1; j < len(strs); j++ {
			strArr := strings.Split(strs[j], "")

			if i >= len(strArr) {
				isMatchAll = false
				continue
			}

			if strArr[i] != str0Arr[i] {
				isMatchAll = false
			}
		}
		if isMatchAll {
			commonStr += str0Arr[i]
		} else {
			break
		}
	}
	return commonStr
}
