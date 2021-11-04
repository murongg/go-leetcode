package bytedance

// 234. 回文链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		if arr[left] != arr[right] {
			return false
		}
	}
	return true
}
