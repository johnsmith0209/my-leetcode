const assert = require('assert')
/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
  let profit = 0
  if (prices.length <= 0) return profit;
  let slowPointer = 0
  let fastPointer = 0
  let currMin = prices[slowPointer]
  let currMax = prices[slowPointer]
  while (slowPointer < prices.length) {
    while (prices[fastPointer + 1] !== undefined && prices[fastPointer + 1] <= currMin) {
      fastPointer++
      currMin = prices[fastPointer]
    }
    slowPointer = fastPointer
    currMax = prices[slowPointer]
    while (prices[fastPointer + 1] !== undefined && prices[fastPointer + 1] >= currMax) {
      fastPointer++
      currMax = prices[fastPointer]
    }
    profit += currMax - currMin
    slowPointer = fastPointer
    currMax = currMin = prices[fastPointer]
    if (prices[fastPointer + 1] === undefined) break;
  }
  return profit
};

assert.equal(maxProfit([7,1,5,3,6,4]), 7)
assert.equal(maxProfit([7,1]), 0)
assert.equal(maxProfit([7,1,6]), 5)
assert.equal(maxProfit([7,1,6,7,2,3,1,5,4,7]), 14)

