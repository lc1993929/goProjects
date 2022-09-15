//  注意：这里一定要注意solc和solidity的版本一定要对应，否则会出现版本不兼容的问题
let solc = require('solc');
let fs = require('fs');

//读取合约
let input = fs.readFileSync('./src/contracts/lottery.sol', 'utf-8');
let output = solc.compile(input, 1);

module.exports=output.contracts[':lottery']