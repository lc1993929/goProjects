// SPDX-License-Identifier: MIT
pragma solidity ^0.4.24;

contract HelloWorld {

    string public message;

    constructor() public {
    }

    function setMessage(string memory msgs) public {
        message = msgs;
    }

    function getMessage() public view returns (string memory){
        return  message;
    }
}
