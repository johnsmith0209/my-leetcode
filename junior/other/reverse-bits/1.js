/**
 * @param {number} n - a positive integer
 * @return {number} - a positive integer
 */
var reverseBits = function(n) {
  n = n.toString(2).split('')
  const newN = []
  const diff = 32 - n.length
  for(let i=0; i<32; i++) {
    if (n[i]) {
      newN[i+diff] = n[i]
    } else {
      newN[i-n.length] = "0"
    }
  }
  n = newN
  for(let i=0; i< 16; i++) {
    const temp = n[i]
    n[i] = n[31-i]
    n[31-i] = temp
  }
  return parseInt(n.join(''), 2)
};