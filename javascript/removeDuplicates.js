const assert = require('assert')

/**
 * https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/21/
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  // console.log(nums)
  if (nums.length <= 1) return nums.length;
  let slowPointer = 0
  let fastPointer = 1
  while (nums[fastPointer] !== undefined) {
    while (nums[fastPointer] === nums[slowPointer]) fastPointer++;
    if (nums[fastPointer] === undefined) break;
    nums[slowPointer + 1] = nums[fastPointer]
    slowPointer++
    fastPointer++
  }
  while (nums.length > slowPointer + 1) nums.pop()
  return slowPointer + 1;
};


assert.equal(removeDuplicates([0,0,0,1,2,3,4]), 5)
assert.equal(removeDuplicates([0,1,2,3,4]), 5)
assert.equal(removeDuplicates([0,1,2,2,3,4]), 5)
assert.equal(removeDuplicates([0,1,2,2,4]), 4)
assert.equal(removeDuplicates([0,1,1,2,2,4]), 4)
assert.equal(removeDuplicates([0,1,1,2,2,3,4,4,5]), 6)