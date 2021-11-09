/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let max = 0
  let length = height.length, left = 0, right = length - 1

  while (left < right) {
    const heightNum = Math.min(height[left], height[right])
    const widthNum = right - left
    const area = widthNum * heightNum
    max = Math.max(area, max)
    if (height[left] <= height[right]) {
      left++
    } else {
      right--
    }
  }
  return max
};


maxArea([1, 1])