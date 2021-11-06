package bytedance

// 还没实现
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length <= 3 {
		return nil
	}
	var res [][]int
	for i := 0; i < length; i++ {
		slow := i + 1
		fast := i + 2
		for fast < length {
			if nums[i]+nums[slow]+nums[fast] == 0 {
				tmp := []int{nums[i], nums[slow], nums[fast]}
				res = append(res, tmp)
				// hashMap[tmp]
			}
			slow++
			fast++
		}
	}

	return res

}
