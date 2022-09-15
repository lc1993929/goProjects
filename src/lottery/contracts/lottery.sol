pragma solidity ^0.4.24;


contract lottery {

    //    管理员地址
    address public manager;
    //    投注人集合
    address[] public players;
    //    期数
    uint256 public round;
    //    中奖人
    address public winner;
    //    每注金额
    uint256 public oneLotteryMoney = 1 ether;

    constructor () public {
        manager = msg.sender;
    }

    //    投注函数
    function play() payable public {
        //        每次只能投注1eth
        require(msg.value == oneLotteryMoney,"每次只能投注1eth");
        players.push(msg.sender);
    }

    //    获取奖池余额
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
    //    listAllPlayers
    function getPlayers() public view returns (address[]){
        return players;
    }

    function getPlayersCount() public view returns (uint){
        return players.length;
    }

    modifier isManager(){
        require(msg.sender == manager, "只有管理员可以开奖");
        _;
    }

    //    开奖
    function openReward() public isManager {
        uint256 random = uint256(keccak256(abi.encodePacked(block.timestamp, block.difficulty, players.length)));
        uint256 index = random % players.length;
        winner = players[index];
        uint256 allMoney = address(this).balance;
        uint256 winnerMoney = allMoney * 90 / 100;
        winner.transfer(winnerMoney);
        manager.transfer(allMoney - winnerMoney);
        round++;
        delete players;

    }
    //    退奖
    function backMoney() public isManager {
        for (uint256 i = 0; i < players.length; i++) {
            players[i].transfer(oneLotteryMoney);
        }

        round++;
        delete players;
    }


}