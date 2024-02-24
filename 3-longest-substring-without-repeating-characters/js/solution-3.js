#! /usr/bin/env node
'use strict';

/**
 * 滑动窗口的题
 *  3. 无重复字符的最长子串
 *  30. 串联所有单词的子串
 *  76. 最小覆盖子串
 *  159. 至多包含两个不同字符的最长子串
 *  209. 长度最小的子数组
 *  239. 滑动窗口最大值
 *  567. 字符串的排列
 *  632. 最小区间
 *  727. 最小窗口子序列
 */

/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    const charSet = new Set();
    let result = 0;
    let startIdx = 0, endIdx = 0;
    while (endIdx < s.length) {
        if (charSet.has(s[endIdx])) {
            charSet.delete(s[startIdx]);
            startIdx++;
        } else {
            charSet.add(s[endIdx]);
            endIdx++;
        }
        result = Math.max(result, endIdx - startIdx);
    }
    return result;
};

console.log(lengthOfLongestSubstring('abcabcbb'));