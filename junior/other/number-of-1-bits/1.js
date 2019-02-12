var hammingWeight = function(n) {
  let rzt = 0
  while(n > 0) {
    rzt += n & 1 === 1 ? 1 : 0
    n >>> 1
  }
  return rzt
};

