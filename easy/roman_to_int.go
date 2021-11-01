package easy

// 13. 罗马数字转整数
func RomanToInt(s string) int {
	numMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	for i := 0; i < len(s); i++ {
		value := numMap[s[i]]
		if i < len(s)-1 && value < numMap[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}
