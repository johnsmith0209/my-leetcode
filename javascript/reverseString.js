/**
 * https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/32/
 * @param {string} s
 * @return {string}
 */
var reverseString = function(s) {
  if (s.length <= 1) return s;
  s = s.split('')
  let start = 0
  let end = s.length - 1
  while (start < end) {
    const tempChar = s[start]
    s[start] = s[end]
    s[end] = tempChar
    start++;
    end--;
  }
  return s.join('')
};
console.log(reverseString('hello'))