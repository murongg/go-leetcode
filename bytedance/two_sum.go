package bytedance

// 1. 两数之和
// 暴力方式
func TwoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hashMap

func TwoSum2(nums []int, target int) []int {
	hashMap := map[int]int{}
	for k, v := range nums {
		if value, ok := hashMap[target-v]; ok {
			return []int{value, k}
		}
		hashMap[v] = k
	}
	return nil
}
