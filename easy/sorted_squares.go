package easy

// 977. 有序数组的平方
func SortedSquares(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	left := 0
	right := length - 1

	for i := right; i >= 0; i-- {
		if L, R := pow(nums[left]), pow(nums[right]); L > R {
			result[i] = L
			left++
		} else {
			result[i] = R
			right--
		}
	}
	return result
}

func pow(num int) int {
	return num * num
}
