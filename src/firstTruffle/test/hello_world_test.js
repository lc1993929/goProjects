const HelloWorldTest = artifacts.require("HelloWorld");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("HelloWorld", function (/* accounts */) {
    it("should assert true", async function () {
        let accounts = await web3.eth.getAccounts()
        let helloWorld = await HelloWorldTest.deployed();
        helloWorld.setMessage("test1", {from: accounts[0]});
        let msg = await helloWorld.getMessage({from: accounts[0]});
        console.log(msg);


        return assert.isTrue(true);
    });
});
