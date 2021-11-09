/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
var reverseString = function (s) {
  let length = s.length, left = 0, right = length - 1

  while (left < right) {
    const tmp = s[left]
    s[left] = s[right]
    s[right] = tmp
    left++
    right--
  }
  return s
};