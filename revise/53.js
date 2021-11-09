/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function (nums) {

  let max = nums[0], pre = 0
  nums.forEach(num => {
    if (pre > 0) {
      pre += num
    } else {
      pre = num
    }

    max = Math.max(pre, max)
  })
  return max
};

maxSubArray([-2, -3, -1, -5])