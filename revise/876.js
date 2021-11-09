/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
// var middleNode = function (head) {

//   const arr = []
//   let tmp = head
//   while (head) {
//     arr.push(head.val)
//     head = head.next
//   }
//   let mid = parseInt(arr.length / 2) + 1
//   for (let i = 0; i < mid - 1; i++) {
//     tmp = tmp.next
//   }
//   return tmp
// };

var middleNode = function (head) {

  const arr = [head]
  while (arr[arr.length - 1].next) {
    arr.push(arr[arr.length - 1].next)
  }

  return arr[Math.trunc(arr.length / 2)]
};

