const assert = require('assert');
/**
 * poor efficiency TLE
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  var rzt = '';
  var chars = [];
  for(var i = 0; i < s.length; i++) {
    chars[0] = s[i];
    for(var j = i + 1; j < s.length; j++) {
      if (chars.indexOf(s[j]) < 0) {
        chars.push(s[j]);
      } else {
        break;
      }
    }
    if (chars.length > rzt.length) {
      rzt = chars.join('');
    }
    chars = []
  }
  return rzt.length;
};

assert.equal(lengthOfLongestSubstring('abcabcbb'), 'abc', 'result of abcabcbb should be abc') 
assert.equal(lengthOfLongestSubstring('abbcabcbb'), 'bca', 'result of abbcabcbb should be bca') 