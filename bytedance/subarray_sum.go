package bytedance

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
