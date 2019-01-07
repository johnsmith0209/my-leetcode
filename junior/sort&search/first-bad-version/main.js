/**
 * Definition for isBadVersion()
 *
 * @param {integer} version number
 * @return {boolean} whether the version is bad
 * isBadVersion = function(version) {
 *     ...
 * };
 */

/**
 * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function(isBadVersion) {
  /**
   * @param {integer} n Total versions
   * @return {integer} The first bad version
   */
  const rztFunc = function(n) {
    let l = 0, r = n, rzt = 0;
    if (n === 0) return 0
    while(true) {
      if (r - l === 1) {
        return isBadVersion(r) ? r : 0;
      }
      const mid = parseInt((r + l) / 2)
      const mBad = isBadVersion(mid)
      console.log(`l[${l}]:${isBadVersion(l)} r[${r}]:${isBadVersion(r)} m[${mid}]:${mBad}`)
      if (mBad) {
        if (mid === 0) {
          return 0
        }
        r = mid
      } else {
        l = mid
      }
    }
  }
  return rztFunc;
};