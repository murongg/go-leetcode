/**
 * @param {number[]} nums
 * @return {number}
 */
var longestConsecutive = function (nums) {
  const set = new Set([...nums])
  let count = 0

  set.forEach(item => {
    if (!set.has(item - 1)) {
      let currVal = item
      let currNum = 1
      while (set.has(currVal + 1)) {
        currNum++
        currVal++
      }
      count = Math.max(currNum, count)
    }
    console.log("------")
  })

  return count
};

longestConsecutive([100, 4, 200, 1, 3, 2])