/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function (head) {
  const map = new Map()
  while (head) {
    if (map.has(head)) {
      return true
    }
    map.set(head)
    head = head.next
  }

  return false
};