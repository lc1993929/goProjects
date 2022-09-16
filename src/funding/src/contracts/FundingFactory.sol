pragma solidity ^0.4.24;

import "./Funding.sol";

contract FundingFactory {

    //    平台管理员
    address platformManager;
    //    所有众筹合约的地址
    address[] allFundings;
    //    创建人自己的合约集合
    mapping(address => address[]) creatorFundings;

    SupportFundingContract supporterFundingsContract;


    constructor() public {
        platformManager = msg.sender;
        supporterFundingsContract = new SupportFundingContract();
    }

    function createFunding(string _projectName, uint256 _targetMoney, uint256 _supportMoney, uint256 _duration) public {
        address fundingAddress = new Funding(msg.sender, _projectName, _targetMoney, _supportMoney, _duration, supporterFundingsContract);
        allFundings.push(fundingAddress);
        creatorFundings[msg.sender].push(fundingAddress);
    }

    function getAllFundings() public view returns (address[]){
        return allFundings;
    }

    function getCreatorFundings() public view returns (address[]){
        return creatorFundings[msg.sender];
    }

    function getSupporterFundings() public view returns (address[]){
        return supporterFundingsContract.getFundings(msg.sender);
    }
}


//因为solidity不允许复杂对象的参数传递，通过这个合约来做不同合约之间的参数传递
contract SupportFundingContract {
    //    参与人的合约集合
    mapping(address => address[]) supporterFundings;

    function addFunding(address _support, address _funding) public {
        supporterFundings[_support].push(_funding);
    }

    function getFundings(address _support) public view returns (address[]){
        return supporterFundings[_support];
    }


}
