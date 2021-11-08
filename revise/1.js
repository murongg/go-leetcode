/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  for (let i = 0; i < nums.length; i++) {
    const left = nums[i];
    for (let j = i + 1; j < nums.length; j++) {
      const right = nums[j];
      if (left + right == target) {
        return [i, j]
      }
    }
  }
  return []
};



var twoSum2 = function (nums, target) {

  const map = new Map()
  for (let i = 0; i < nums.length; i++) {
    if (map.has(target - nums[i])) {
      return [map.get(target - nums[i]), i]
    }
    map.set(nums[i], i)
  }

  return []

}

