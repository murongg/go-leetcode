package bytedance

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
