package easy

// 二分查找
func TwoSum1(numbers []int, target int) []int {
	length := len(numbers)

	for i := 0; i < length; i++ {
		left := i + 1
		right := length - 1
		for left <= right {
			mid := left + ((right - left) / 2)
			if numbers[mid]+numbers[i] == target {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return []int{}
}

// 双指针
func TwoSum2(numbers []int, target int) []int {
	length := len(numbers)
	left := 0
	right := length - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
