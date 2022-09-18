const HelloWorld = artifacts.require("HelloWorld");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("HelloWorld", function (/* accounts */) {
    it("should assert true", async function () {
        await HelloWorld.deployed();
        return assert.isTrue(true);
    });
});


//与合约交互
// var contract = require("truffle-contract");
// var MyContract = contract(SimpleStorage.json的内容)
// MyContract.setProvider()
// var deployedInstance;
// MyContract.deployed().then(function (instance) {
//     var deployedInstance = instance;
//
//
//     let accounts = await web3.eth.getAccounts();
//     deployedInstance.sendCoin(accounts[1], 10, {from: accounts[0]});
//     deployedInstance.getValue.call(参数填这⾥, {from:xxx}).then(res)
//
// });

