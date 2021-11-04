package bytedance

/**
 * 206. 反转链表
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	// 保存个空节点，这是即将返回的节点
	var prev *ListNode
	// 当前节点
	curr := head

	for curr != nil {
		// 保存下一个节点
		next := curr.Next
		// 让当前节点的下一个节点，指向上一个节点
		curr.Next = prev
		// 让上一个节点指向当前节点
		prev = curr
		// 让当前节点指向下个节点
		curr = next
	}
	return prev
}
