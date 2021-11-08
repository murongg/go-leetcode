/**
 * @param {string} s
 * @return {number}
 */

const map = {
  "I": 1,
  "V": 5,
  "X": 10,
  "L": 50,
  "C": 100,
  "D": 500,
  "M": 1000
}
var romanToInt = function (s) {
  let res = 0

  for (let i = 0; i < s.length; i++) {

    const value = map[s[i]]
    if (value < map[s[i + 1]]) {
      res -= value
    } else {
      res += value
    }
  }
  console.log(res)
  return res
};



romanToInt("MCMXCIV")