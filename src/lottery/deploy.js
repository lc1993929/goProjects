let {bytecode, interface: inter} = require('./compile');

const Web3 = require('web3');
let HDWalletProvider = require('@truffle/hdwallet-provider');
//助记词
let fs = require('fs');
//读取文件
const memoryWords = fs.readFileSync('F:\\liuchang\\memoryWords.txt', 'utf-8');
let provider = new HDWalletProvider({
    mnemonic: {
        phrase: memoryWords
    },
    providerOrUrl: 'HTTP://127.0.0.1:7545',
    // providerOrUrl: 'https://goerli.infura.io/v3/9b08343587a3477f9be3a21525d7baaf'
    // providerOrUrl: 'https://sepolia.infura.io/v3/9b08343587a3477f9be3a21525d7baaf'
});

const web3 = new Web3(provider);
let contract = new web3.eth.Contract(JSON.parse(inter));


let doDeploy = async () => {
    //获取账户
    let accounts = await web3.eth.getAccounts();
    // console.log(accounts);
    let account = accounts[0];

    let newContractInstance = await contract.deploy({data: bytecode,}).send({
        from: account,
        gas: 1500000,
        // gasPrice: '3000000'
    });
    console.log(newContractInstance.options.address)
}

doDeploy().then();