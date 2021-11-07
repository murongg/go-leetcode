package bytedance

// 贪心算法

// 55. 跳跃游戏
func CanJump(nums []int) bool {

	cur := 0

	for i := 0; i < len(nums); i++ {
		if i > cur {
			return false
		}
		if nums[i]+i > cur {
			cur = i + nums[i]
		}
	}

	return true
}

// [2,3,1,1,4]

// index value nextIndex
//   0     1      2
//   2     3      5
// [3,2,1,0,4]
// index value nextIndex
//   0     3     3
//   3     0     false

// num[length - 1] == nums[]

// 45. 跳跃游戏 II
func Jump(nums []int) int {

	step, end, maxPos := 0, 0, 0
	i := 0
	for end < len(nums)-1 {
		if nums[i]+i > maxPos {
			maxPos = nums[i] + i
		}
		if end == i {
			end = maxPos
			step++
		}
		i++
	}
	return step
}
