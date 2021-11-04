package bytedance

// 141. 环形链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢指针
// 快指针在初始化时在满指针前面
// 快指针每次移动两步，慢指针移动一步
// 当快指针与慢指针相遇，则是环形，否则不是
// 一直循环到快指针为nil，并且快指针的 next 也为 nil 时结束
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 利用 hashMap
// 每次去hashMap里寻找，如果找到，说明相遇，则返回 true
// 如果找不到继续往next找
// 直到 head 为空
func HasCycleHashMap(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	hashMap := map[*ListNode]struct{}{}

	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		}
		hashMap[head] = struct{}{}
		head = head.Next
	}
	return false
}
