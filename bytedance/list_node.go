package bytedance

type ListNode struct {
	Val  int
	Next *ListNode
}

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

// 160. 相交链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  hashMap解法
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := map[*ListNode]struct{}{}

	for headA != nil {
		hashMap[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

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

// 合并链表
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

/**
 * 206. 反转链表
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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
