package bytedance

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
