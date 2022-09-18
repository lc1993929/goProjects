pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/HelloWorld.sol";


contract TestHelloWorld {
    HelloWorld instance = HelloWorld(DeployedAddresses.HelloWorld());

    function testHelloWorld() public {

        string memory text = "Hello World";

        instance.setMessage(text);
        string memory nowMsg = instance.getMessage();


        Assert.equal(nowMsg, text, "msg error");
    }
}
