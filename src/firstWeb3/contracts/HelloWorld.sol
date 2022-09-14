pragma solidity ^0.4.24;

contract HelloWorld {


    string public message;
    uint public value;
    bytes public chars;

    enum WeekDays {
        Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, sunday
    }

    constructor(string _initMsg) public {
        message = _initMsg;
    }

    // function Inbox() payable public {}



    function setMessage(string newMessage) public {
        message = newMessage;
    }

    function getMessage() public view returns (string){
        return message;
    }

    function test1() public payable {
        value = msg.value;
        chars = "testsd";
        bytes(message).length;
    }


}