/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function (nums, k) {

  let res = 0

  for (let i = 0; i < nums.length; i++) {

    let tmp = nums[i]

    for (let j = i + 1; j < nums.length && tmp < k; j++) {

      tmp += nums[j]
    }
    console.log(tmp)


  }
};

subarraySum([1, 2, 1], 2)