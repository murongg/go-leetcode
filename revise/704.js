/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function (nums, target) {
  let left = 0, right = nums.length - 1

  while (left <= right) {
    const mid = Math.floor((right - left) / 2) + left;
    if (nums[mid] === target) {
      return mid
    }
    if (nums[left] < target) {
      left++
    }

    if (nums[right] > target) {
      right--
    }
  }
  return -1
};


console.log(search([5], 5)
)