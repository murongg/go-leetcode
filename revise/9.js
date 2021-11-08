/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  if (x < 0) {
    return false
  }
  let tmp = x
  let res = 0

  while (x !== 0) {
    res = res * 10 + x % 10
    x = ~~(x / 10)
  }
  return res === tmp
};

console.log(isPalindrome(121))