// https://leetcode-cn.com/problems/number-of-digit-one/description/
const assert = require('assert')
/**
 * @param {number} n
 * @return {number}
 */
var countDigitOne = function(n) {
  if (n <= 0) return 0
  const maxDigit = `${n}`.length
  if (maxDigit === 1) return n === 0 ? 0 : 1;
  const digits = Math.pow(10, maxDigit - 1)
  let times = Math.floor(n / digits)
  const reminder = n % Math.pow(10, maxDigit - 1)
  let lowerRemain = n - reminder
  let firstPart = countDigitOne(reminder),
      secondPart = 0;
  if (times === 1) {
    firstPart += 1 + reminder
  }
  if (lowerRemain / digits > 2) {
    secondPart = (times - 2) * countDigitOne((lowerRemain - 1) % digits)
    secondPart += countDigitOne((lowerRemain - 1) % digits + digits)
  } else {
    secondPart += countDigitOne(lowerRemain - 1)
  }
  return firstPart + secondPart
};

assert.equal(countDigitOne(12),5)
assert.equal(countDigitOne(13),6)
assert.equal(countDigitOne(14),7)
assert.equal(countDigitOne(4),1)
assert.equal(countDigitOne(19),12)
assert.equal(countDigitOne(21),13)
assert.equal(countDigitOne(23),13)
assert.equal(countDigitOne(34),14)
assert.equal(countDigitOne(41),15)
assert.equal(countDigitOne(45),15)
assert.equal(countDigitOne(126),60)
assert.equal(countDigitOne(123),57)
assert.equal(countDigitOne(235),154)
assert.equal(countDigitOne(512),205)
assert.equal(countDigitOne(1241),697)
assert.equal(countDigitOne(5126),2560)
assert.equal(countDigitOne(31414),22902)