/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
  let res = 0
  while (x !== 0) {
    const tmp = x % 10
    x = parseInt(x / 10)
    res = res * 10 + tmp
    // 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
    if (res < Math.pow(-2, 31) || res > Math.pow(2, 31) - 1) {
      return 0;
    }
  }

  return res
};


reverse(-321)