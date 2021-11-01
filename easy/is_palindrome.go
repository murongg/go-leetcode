package easy

import "math"

/**
	题目：
	9. 回文数
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
**/

func IsPalindrome(x int) bool {
	tmp := x
	if x < 0 {
		return false
	}

	var result int
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return false
		}
		last := int(x % 10)
		x = x / 10
		result = result*10 + last
	}

	return result == tmp
}
