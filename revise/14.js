/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {

  let prev = strs[0]
  for (let i = 1; i < strs.length; i++) {
    const prefix = getPrefix(strs[i], prev)
    prev = prefix
  }

  return prev
};


function getPrefix(str1, str2) {
  let str = ""
  if (str1.length < str2.length) {
    str = str1
  } else {
    str = str2
  }

  let res = ""

  for (let i = 0; i < str.length; i++) {
    if (str1[i] === str2[i]) {
      res = res + str1[i]
    } else {
      return res
    }
  }
  return res
}

longestCommonPrefix(["flower", "flow", "flight"])