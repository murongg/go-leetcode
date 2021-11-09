/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (numbers, target) {
  let length = numbers.length, left = 0, right = length - 1

  while (left < right) {
    const l = numbers[left]
    const r = numbers[right]
    if (l + r === target) {
      return [left + 1, right + 1]
    } else if (l + r > target) {
      right--
    } else {
      left++
    }
  }

  return null
};


console.log(twoSum([5, 25, 75], 100))