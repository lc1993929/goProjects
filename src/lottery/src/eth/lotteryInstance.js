let abiFile = require('../contracts/artifacts/lottery_metadata.json');

// const address = '0xdE7F6aa3182cC40189D92B069A85Fd968b728ADc';
const address = '0xAE1B676DC34e63aeEeE58ebB6791065907F35907';
let web3 = require('../utils/initWeb3');
let abi = abiFile.output.abi;

let contract = new web3.eth.Contract(abi, address);
// console.log(contract.options.address);

module.exports = contract;