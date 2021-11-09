/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function (nums) {
  let left = 0, length = nums.length, right = length - 1, result = new Array(length)

  for (let i = right; i >= 0; i--) {
    const l = Math.pow(nums[left], 2)
    const r = Math.pow(nums[right], 2)
    if (l < r) {
      result[i] = r
      right--
    } else {
      result[i] = l
      left++
    }
  }
  return result
};
sortedSquares([-4, -1, 0, 3, 10])