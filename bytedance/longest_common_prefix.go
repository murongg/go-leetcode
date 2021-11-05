package bytedance

func LongestCommonPrefix(strs []string) string {
	// 如果数组长度等于0，则绝对没有公共前缀
	if len(strs) == 0 {
		return ""
	}
	// 先获取数组第一个字符串
	first := strs[0]
	// 前缀
	var prefix string
	// 从下标 1 开始循环 strs
	for i := 1; i < len(strs); i++ {
		// 查找两者公共前缀
		prefix = lcp(first, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(str1 string, str2 string) string {
	// 获取两者的最小长度，遍历使用
	length := min(str1, str2)
	// 默认下标
	index := 0
	// 循环 index < length && 两字符串同一下标的值相等时，则存在公共字符串，继续 index++，直到获取不到相等数据
	for index < length && str1[index] == str2[index] {
		index++
	}
	// 获取str1，str2任意一个数据的index之前的字符串
	return str1[index:]
}

func min(str1 string, str2 string) int {
	if len(str1) > len(str2) {
		return len(str2)
	} else {
		return len(str1)
	}
}
