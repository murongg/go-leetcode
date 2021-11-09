/**
 * @param {number[]} nums
 * @return {number}
 */
var findLengthOfLCIS = function (nums) {
  let max = 1, tmp = 1
  let slow = 0, fast = 1

  while (fast < nums.length) {
    if (nums[slow] < nums[fast]) {
      tmp++
      max = Math.max(max, tmp)
    } else {
      tmp = 1
    }
    slow++
    fast++
  }
  return max
};


findLengthOfLCIS([2, 3, 4])