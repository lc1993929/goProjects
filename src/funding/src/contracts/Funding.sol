pragma solidity ^0.4.24;

import "./FundingFactory.sol";

contract Funding {

    address public manager;
    string public projectName;
    uint256 public targetMoney;
    uint256 public supportMoney;
    uint256 public endTime;
    //投资人的集合
    address[] public supporters;
    // 为了方便判断是否有投资人
    mapping(address => bool) supportersMap;
    //    请求集合
    Request[] requests;

    SupportFundingContract supporterFundingsContract;


    constructor (address _creator, string _projectName, uint256 _targetMoney, uint256 _supportMoney, uint256 _duration, SupportFundingContract _supporterFundingsContract) public {
        manager = _creator;
        projectName = _projectName;
        targetMoney = _targetMoney;
        supportMoney = _supportMoney;
        endTime = block.timestamp + _duration;
        supporterFundingsContract = _supporterFundingsContract;
    }

    modifier isManager(){
        require(msg.sender == manager, "非项目管理人");
        _;
    }

    //    支持的方法
    function support() payable public {
        require(msg.value == supportMoney, "投资金额错误");
        supporters.push(msg.sender);
        supportersMap[msg.sender] = true;

        supporterFundingsContract.addFunding(msg.sender, this);
    }

    //    退款
    function refund() public isManager {
        for (uint256 i = 0; i < supporters.length; i++) {
            supporters[i].transfer(supportMoney);
        }
        delete (supporters);
    }

    function getBalance() public view returns (uint256){
        return address(this).balance;
    }

    function getSupporters() public view returns (address[]){
        return supporters;
    }

    //        当前状态。0：投票中，1：已批准，2：已完成，3：拒绝
    enum RequestStatus {
        Voting, Approved, Completed, Reject
    }

    struct Request {
        string purpose;
        uint256 cost;
        address shopAddress;
        uint256 voteCount;
        // 记录已经投票的地址，方便判断是否已经投过票了
        mapping(address => bool) votedMap;

        RequestStatus status;
    }

    //    新增请求
    function createRequest(string _purpose, uint256 _cost, address _shopAddress) public isManager {
        Request memory request = Request({
        purpose : _purpose,
        cost : _cost,
        shopAddress : _shopAddress,
        voteCount : 0,
        status : RequestStatus.Voting
        });

        requests.push(request);
    }

    // 批转请求
    function approveRequest(uint256 _requestId) public {
        // 先从集合中拿到请求
        Request storage request = requests[_requestId];
        // 查看是否有投票权
        require(supportersMap[msg.sender], "只有投资人有投票权");
        // 查看是否已经投过票了
        require(request.votedMap[msg.sender] == false, "已经投过票了");

        // 验证成功，开始投票
        request.votedMap[msg.sender] = true;
        request.voteCount++;

        // 验证赞成人数过半修改为已批准状态
        if (request.voteCount > supporters.length / 2) {
            request.status = RequestStatus.Approved;
        }


    }

    // 执行请求
    function finalizeRequest(uint256 _requestId) public isManager {
        // 先从集合中拿到请求
        Request storage request = requests[_requestId];
        // 验证合约中金额是否足够
        require(address(this).balance >= request.cost, "合约中金额不足");
        // 验证赞成人数过半
        require(request.voteCount > supporters.length / 2, "赞成人数未过半");
        // 转账给商家地址
        request.shopAddress.transfer(request.cost);
        // 修改状态
        request.status = RequestStatus.Completed;

    }

    function getSupportersCount() public view returns (uint256){
        return supporters.length;
    }

    //    剩余时间
    function getRemainTime() public view returns (uint256){
        if (endTime <= now) {
            return 0;
        }
        return (endTime - now) / 60 / 60 / 24;
        
    }

    function getRequestsCount() public view returns (uint256){
        return requests.length;
    }


    function getRequestDetailByIndex(uint256 _requestId) public view returns (string, uint256, address, uint256, bool, uint256){
        require(_requestId < requests.length, "未知请求id");
        Request storage request = requests[_requestId];
        bool isVoted = request.votedMap[msg.sender];
        return (request.purpose, request.cost, request.shopAddress, request.voteCount, isVoted, uint256(request.status));

    }


}
