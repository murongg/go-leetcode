package bytedance

import "fmt"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	var prefix string
	for i := 1; i < len(strs); i++ {
		prefix = lcp(first, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(str1 string, str2 string) string {
	length := min(str1, str2)
	index := 0
	fmt.Println(length)
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[index:]
}

func min(str1 string, str2 string) int {
	if len(str1) > len(str2) {
		return len(str2)
	} else {
		return len(str1)
	}
}
