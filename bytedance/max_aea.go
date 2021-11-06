package bytedance

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
