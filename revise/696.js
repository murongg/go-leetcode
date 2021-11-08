/**
 * @param {string} s
 * @return {number}
 */
var countBinarySubstrings = function (s) {
  let pre = 0, last = 0, ans = 0

  const length = s.length
  while (pre < length) {
    let count = 0
    const curr = s[pre]
    while (pre < length && s[pre] === curr) {
      pre++
      count++
    }
    ans += Math.min(last, count)
    last = count
  }

  return ans
};

countBinarySubstrings("00111011")

// 我们可以将字符串 ss 按照 00 和 11 的连续段分组，存在 {counts} 数组中，
// 例如 s=00111011，可以得到这样的 {counts} 数组：counts={2,3,1,2}。

