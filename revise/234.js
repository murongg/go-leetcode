/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function (head) {

  const arr = []
  while (head !== null) {
    arr.push(head.val)
    head = head.next
  }

  let length = arr.length, left = 0, right = length - 1

  while (left < right) {
    if (arr[left] !== arr[right]) {
      return false
    }
    left++
    right--
  }

  return true

};

