//  注意：这里一定要注意solc和solidity的版本一定要对应，否则会出现版本不兼容的问题
let solc = require('solc');
let fs = require('fs');

//读取合约
let input = fs.readFileSync('./contracts/HelloWorld.sol', 'utf-8');
// console.log(input);

let output = solc.compile(input, 1);
// console.log(output);
// console.log(output.contracts[':HelloWorld']);
// for (let contractName in output.contracts) {
//     code and ABI that are needed by web3
    // console.log(contractName + ': ' + output.contracts[contractName].bytecode)
    // console.log(contractName + '; ' + JSON.parse(output.contracts[contractName].interface))
// }


// console.log(output.contracts[':HelloWorld']['interface']);

module.exports=output.contracts[':HelloWorld']