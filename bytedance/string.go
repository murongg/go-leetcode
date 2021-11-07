package bytedance

// 3. 无重复字符的最长子串
// "abcabcbb"
func LengthOfLongestSubstring(s string) int {
	sum := 0
	left := -1
	n := len(s)

	hashMap := map[byte]bool{}
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(hashMap, s[i-1])
		}
		for left+1 < n && !hashMap[s[left+1]] {
			hashMap[s[left+1]] = true
			left++
		}
		currentCount := left - i + 1
		// left-i+1代表一共向右走的个数
		if currentCount > sum {
			sum = currentCount
		}
	}
	return sum
}

// 392. 判断子序列

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
