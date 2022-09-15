const Web3 = require('web3');

const web3 = new Web3(window.ethereum);
// console.log(web3.version);
// console.log(web3.eth.net.currentProvider);

module.exports = web3;