import React from "react";
import CardExampleCard from "./display/ui";

let web3 = require('./utils/initWeb3');

let lotteryInstance = require('./eth/lotteryInstance');

window.ethereum.on('accountsChanged', (accounts) => window.location.reload());

export default class APP extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            manager: '',
            winner: '',
            round: 0,
            playersCount: 0,
            balance: 0,
            players: [],
            currentAccount: '',
            isClicked: false,
            isShowButton: 'none',

        };


    }


    async componentDidMount() {
        let manager = await lotteryInstance.methods.manager().call();
        let round = await lotteryInstance.methods.round().call();
        let winner = await lotteryInstance.methods.winner().call();
        let playersCount = await lotteryInstance.methods.getPlayersCount().call();
        let balance = await lotteryInstance.methods.getBalance().call();
        balance = web3.utils.fromWei(balance);
        let players = await lotteryInstance.methods.getPlayers().call();
        let accounts = await web3.eth.getAccounts();
        let currentAccount = accounts[0];
        let isShowButton = currentAccount === manager ? 'inline' : 'none';

        this.setState({
            manager, winner, round, playersCount, balance, players, currentAccount, isShowButton
        });

    }


    play = async () => {
        this.setState({
            isClicked: true
        });

        try {
            await lotteryInstance.methods.play().send({
                from: this.state.currentAccount,
                value: web3.utils.toWei('0.000001'),
                gas: '2000000'
            });
            alert("投注成功");

        } catch (e) {
            console.log(e);
            alert("投注失败");
        } finally {
            this.setState({
                isClicked: false
            });
            this.componentDidMount();
        }
    }

    openReward = async () => {
        this.setState({
            isClicked: true
        });

        try {
            await lotteryInstance.methods.openReward().send({
                from: this.state.currentAccount,
                gas: '2000000'
            });

            let winner = await lotteryInstance.methods.winner().call();
            alert(`开奖成功！\n 中奖人：${winner}`);

        } catch (e) {
            alert("开奖失败");
        } finally {
            this.setState({
                isClicked: false
            });
            this.componentDidMount();
        }
    }

    backMoney = async () => {
        this.setState({
            isClicked: true
        });

        try {
            await lotteryInstance.methods.backMoney().send({
                from: this.state.currentAccount,
                gas: '2000000'
            });
            alert("退奖成功");

        } catch (e) {
            alert("退奖失败");
        } finally {
            this.setState({
                isClicked: false
            });
            this.componentDidMount();
        }
    }

    render() {
        return (<div className="App">
            <link
                async
                rel="stylesheet"
                href="https://cdn.jsdelivr.net/npm/semantic-ui@2/dist/semantic.min.css"
            />
            <header className="App-header">
                <CardExampleCard manager={this.state.manager} currentAccount={this.state.currentAccount}
                                 playersCount={this.state.playersCount} balance={this.state.balance}
                                 round={this.state.round} play={this.play} openReward={this.openReward}
                                 backMoney={this.backMoney} isClicked={this.state.isClicked}
                                 isShowButton={this.state.isShowButton} winner={this.state.winner}/>

                <p>manager:{this.state.manager}</p>
                <p>winner:{this.state.winner}</p>
                <p>round:{this.state.round}</p>
                <p>playersCount:{this.state.playersCount}</p>
                <p>balance:{this.state.balance}</p>
                <p>players:{this.state.players}</p>
                <p>currentAccount:{this.state.currentAccount}</p>
            </header>
        </div>);
    }


}