package bytedance

// 128. 最长连续序列
// 使用hashMap
func LongestConsecutive(nums []int) int {
	hashMap := map[int]bool{}

	// 利用 Map 去重
	for _, v := range nums {
		hashMap[v] = true
	}

	// 结果
	res := 0

	// 遍历循环后的 Map
	for k := range hashMap {
		// 使用 hashMap[k-1] 是否有连续
		if !hashMap[k-1] {
			// 当前值，用来查找是否有连续值
			currentVal := k
			// 连续值个数，1 开始
			currentNum := 1
			// 遍历 hashMap ，如果 currentVal + 1 存在，则存在连续，
			// currentVal++ 继续 + 1遍历，查找更多连续的个数
			for hashMap[currentVal+1] {
				currentNum++
				currentVal++
			}

			// 如果当前 k 的连续个数大于 res 个数，则 res = currentNum
			if currentNum > res {
				res = currentNum
			}
		}
	}

	return res
}
