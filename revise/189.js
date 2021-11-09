/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function (nums, k) {
  k %= nums.length
  nums = reverse(nums, 0, nums.length - 1)
  nums = reverse(nums, 0, k - 1)
  nums = reverse(nums, k, nums.length - 1)

  return nums

};

function reverse(arrs, start, end) {

  while (start < end) {
    const tmp = arrs[start]
    arrs[start] = arrs[end]
    arrs[end] = tmp
    start++
    end--
  }
  return arrs
}

rotate([1, 2, 3, 4, 5, 6, 7], 3)