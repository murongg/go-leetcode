/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let slow = 1, fast = 1

  while (fast < nums.length) {
    if (nums[fast] !== nums[fast - 1]) {
      nums[slow] = nums[fast]
      slow++
    }
    fast++
  }

  return nums.length
};
removeDuplicates([0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 4])
