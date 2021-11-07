package bytedance

// 70. 爬楼梯
func ClimbStairs(n int) int {

	if n < 2 {
		return 1
	}
	// n-2
	first := 0
	// n-1
	second := 1
	for i := 0; i < n; i++ {
		second, first = first+second, second
	}
	return second
}

// 3 层楼梯
// 1 + 1 + 1
// 1+2
// 2+1

// 2层楼梯
// 1+1
// 2

// f(n) = f(n-1) + f(n-2)
