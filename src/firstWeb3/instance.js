// const Web3 = require('web3');
// const web3 = new Web3(Web3.givenProvider || "ws://localhost:8545");
// const web3 = new Web3("HTTP://127.0.0.1:7545");

let HDWalletProvider = require('@truffle/hdwallet-provider');
//助记词
let fs = require('fs');
const Web3 = require("web3");
//读取文件
const memoryWords = fs.readFileSync('F:\\liuchang\\memoryWords.txt', 'utf-8');
let provider = new HDWalletProvider({
    mnemonic: {
        phrase: memoryWords
    },
    // providerOrUrl: 'https://goerli.infura.io/v3/9b08343587a3477f9be3a21525d7baaf'
    providerOrUrl: 'https://sepolia.infura.io/v3/9b08343587a3477f9be3a21525d7baaf'
});
const web3 = new Web3(provider);

let abi = [
    {
        "constant": true,
        "inputs": [],
        "name": "chars",
        "outputs": [
            {
                "name": "",
                "type": "bytes"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "newMessage",
                "type": "string"
            }
        ],
        "name": "setMessage",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "value",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [],
        "name": "test1",
        "outputs": [],
        "payable": true,
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "getMessage",
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "message",
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "name": "_initMsg",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "constructor"
    }
]
const address = '0x13a8725f126b47ce0de121225776c5f3f5c3f805';

let contract = new web3.eth.Contract(abi, address);

console.log(contract.options.address);
module.exports = contract;