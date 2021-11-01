package easy

import "math"

/**
	题目：
	7. 整数反转
	给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
 	如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
 	假设环境不允许存储 64 位整数（有符号或无符号）。
**/

/**
	示例：
	输入：x = 123 输出：321
	输入：x = -123 输出：-321
	输入：x = 120 输出：21
	输入：x = 0 输出：0
**/

func Reverse(x int) int {
	var result int
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return 0
		}
		last := int(x % 10)
		x = x / 10
		result = result*10 + last
	}
	return result
}
