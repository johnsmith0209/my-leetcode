const assert = require('assert');
// this failed due to not considering substring inside between duplicated chars
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  var rzt = '';
  var charPos = {}; // {char: [idx...]} keep the idx array length <=2
  var longest = {length: 0, char: '', start: 0, end: 0}; // {char: char, width: num}
  for(var i = 0; i< s.length; i++) {
    var curChar = s[i];
    if (!charPos[curChar]) {
      charPos[curChar] = [];
    }
    var curCharIdxArray = charPos[curChar];
    curCharIdxArray.push(i);
  }
  for(var char in charPos) {
    // only once
    var length = 0;
    var start = charPos[char][0];
    var end = charPos[char][1] || charPos[char][0];
    if (charPos[char].length === 1) {
      length = s.length - charPos[char][0];
    } else {
      length = charPos[char][1] - charPos[char][0];
      for(var insideChar in charPos) {
        if (insideChar === char || charPos[insideChar].length === 1) {
          continue;
        }
        for(var i = 0; i < charPos[insideChar].length - 1; i++) {
          var insideBegin = charPos[insideChar][i];
          var insideEnd = charPos[insideChar][i+1];
          console.log('inside', insideChar, insideBegin, insideEnd)
          if (insideBegin > start && insideEnd < end) {
            console.log('inside overlap', insideChar, insideBegin, insideEnd)            
            var curGap = start - insideBegin;
            if (curGap < length) {
              length = curGap;
              end = insideBegin;
            }
          }
        }
      }
    }
    if (length > longest.length) {
      longest = {
        char: char,
        length: length,
        start: start,
        end: end
      }
    }
  }
  console.log(s, charPos, longest);
  return s.slice(longest.start, longest.end);
};

console.log(lengthOfLongestSubstring('abcabcbb'))
console.log(lengthOfLongestSubstring('abbcabcbb'))
assert.equal(lengthOfLongestSubstring('abcabcbb'), 'abc', 'result of abcabcbb should be abc') 
assert.equal(lengthOfLongestSubstring('abbcabcbb'), 'bca', 'result of abbcabcbb should be bca') 