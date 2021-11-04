package bytedance

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟链表
	dummy := new(ListNode)
	// 当前节点，指针指向头部
	curr := dummy
	// 只要是 l1或l2 不为空，就一直循环
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			// 当 l1 <= l2 时，指针移动到l1
			curr.Next = l1
			// l1 移动到 l1 的下一级
			l1 = l1.Next
		} else {
			// 当 l1 > l2 时，指针移动到l2
			curr.Next = l2
			// l2 移动到 l2 的下一级
			l2 = l2.Next
		}
		// 指针向后移动
		curr = curr.Next
	}

	// 处理 l1 最后剩余数据
	if l1 != nil {
		curr.Next = l1
	}
	// 处理 l2 最后剩余数据
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}
