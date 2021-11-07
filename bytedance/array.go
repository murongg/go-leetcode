package bytedance

// 最长公公前缀
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

// 674. 最长连续递增序列
// 思路：快慢指针
func FindLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 1
	result, tmp := 1, 1

	for fast < len(nums) {
		if nums[slow] < nums[fast] {
			tmp++
			if tmp > result {
				result = tmp
			}
		} else {
			tmp = 1
		}
		slow++
		fast++
	}
	return result
}

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

// 11. 盛最多水的容器
func MaxArea(height []int) int {

	left := 0
	right := len(height) - 1
	max := 0
	for left < right {
		tmp := 0
		if height[left] > height[right] {
			tmp = (right - left) * height[right]
			right--
		} else {
			tmp = (right - left) * height[left]
			left++
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// 53. 最大子序和
// 动态规划
// 若当前元素大于0，则将其加到当前元素上
func MaxSubArray(nums []int) int {
	max := nums[0]
	pre := 0

	for _, v := range nums {
		if pre > 0 {
			pre = pre + v
		} else {
			pre = v
		}
		if pre > max {
			max = pre
		}
	}
	return max
}

// 26. 删除有序数组中的重复项
// 快慢指针
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	fast := 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 560. 和为 K 的子数组
// 暴力破解，每次计算start之前的所有数组
// 比如[1,-1,0] k=0
// start 为 0 时，计算 [1] = 1：0+1=0
// start 为 1 时，计算 [1,-1]： -1+0=-1，-1+1=0
// start 为 2 时，计算 [1,-1,0] = 0+0=0，0-1=-1，0-1+1=0
func SubarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// hashMap

func SubarraySumHashMap(nums []int, k int) int {
	count, pre := 0, 0
	hashMap := map[int]int{}
	hashMap[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := hashMap[pre-k]; ok {
			count += hashMap[pre-k]
		}
		hashMap[pre] += 1
	}
	return count
}
