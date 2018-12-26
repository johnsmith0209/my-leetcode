/**
 * @param {number} x
 * @param {number} y
 * @return {number}
 */
var hammingDistance = function(x, y) {
  return (x^y).toString(2).split("").reduce((memo, item) => {return +item === 1 ? memo+1:memo}, 0)
};