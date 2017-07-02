var _ = require('lodash');

var str = [
  'a', 'b', 'c', 'd', 'e', 'f', 'g',
  'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
  'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'
];

var genStr = (len) => {
  var str = '';
  for(var i = 0; i < len; i++) {
    str += genC();
  }
  return str;
};

var genC = () => {
  return str[genRN()];
};

var genRN = (max = 24) => {
  return parseInt(Math.random() * max);
}

var genT = (str) => {
  var c = genC();
  var pos = genRN(str.length);
  var rzt = str;
  console.log(`gen pos: ${pos}, c: ${c}`)
  rzt = str.slice(0, pos) + c + str.slice(pos);
  // return rzt;
  return _.shuffle(rzt).join('');
};

module.exports = {
  genStr, genC, genRN, str, genT
}