package bytedance

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
