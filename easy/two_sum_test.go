package easy_test

import (
	"leetcode/easy"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	success := []int{0, 1}
	result := easy.TwoSum(nums, target)

	if success[0] == result[0] && success[1] == result[1] {
		t.Log("success")
	}
}
