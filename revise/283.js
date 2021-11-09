/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function (nums) {
  let left = 0, right = 0, length = nums.length
  while (right < length) {
    if (nums[right] != 0) {
      const tmp = nums[left]
      nums[left] = nums[right]
      nums[right] = tmp
      left++
    }

    right++
  }
  return nums
};

console.log(moveZeroes([0, 0, 1]))