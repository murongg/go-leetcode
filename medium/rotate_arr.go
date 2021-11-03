package medium

// 189. 旋转数组
// 先将整个数组反转
// 再反转两边
func reverse(nums []int) {
	length := len(nums)
	right := length - 1
	for left := 0; left < length/2; left++ {
		nums[left], nums[right] = nums[right], nums[left]
		right--
	}
}

func RotateArr(nums []int, k int) {
	// 防止边界溢出
	k %= len(nums)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}
