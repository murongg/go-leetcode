
function LongestCommonPrefix(strs) {
  if (strs.length === 0) {
    return ""
  }
  let first = strs[0]
  let prefix = ""
  for (let i = 1; i < strs.length; i++) {
    prefix = lcp(first, strs[i])
    if (prefix.length === 0) {
      return ""
    }
  }
  return prefix
}

function lcp(str1, str2) {
  const len = Math.min(str1.length, str2.length)
  let index = 0
  while (index < len && str1[index] === str2[index]) {
    index++
  }
  return str1.substr(0, index)
}
