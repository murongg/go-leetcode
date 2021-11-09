/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  let oldV = 0, newV = 1
  for (let i = 0; i < n; i++) {
    const tmp = newV
    newV = oldV + newV
    oldV = tmp
    // newV, oldV = oldV + newV, newV
  }
  console.log(newV)
  return newV
};

climbStairs(10)