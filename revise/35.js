/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function (nums, target) {
  // let left = 0, right = nums.length - 1

  // while (left <= right) {
  //   var mid = Math.floor(left + ((right - left) / 2))
  //   if (nums[mid] === target) {
  //     return mid
  //   } else if (right > target) {
  //     right = mid - 1
  //   } else {
  //     left = mid + 1
  //   }
  // }
  // return left
  let left = 0, right = nums.length - 1

  while (left <= right) {
    var mid = Math.floor(left + ((right - left) / 2))
    if (nums[mid] === target) {
      return mid
    } else if (nums[mid] > target) {
      right = mid - 1
    } else {
      left = mid + 1
    }
  }
  return left
};


console.log(searchInsert([1, 3, 5, 6], 5))