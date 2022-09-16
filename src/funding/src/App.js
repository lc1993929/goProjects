import './App.css';
import React from 'react';
import TabCenter from "./display/tabCenter";
import web3 from './utils/initWeb3';

window.ethereum.on('accountsChanged', (accounts) => window.location.reload());

class App extends React.Component {
    constructor() {
        super();

        this.state = {
            currentAccount: ''
        }
    }

    async componentDidMount() {
        let accounts = await web3.eth.getAccounts();
        this.setState({
            currentAccount: accounts[0]
        });

    }


    render() {
        return (


            <div className="App">
                <link
                    async
                    rel="stylesheet"
                    href="https://cdn.jsdelivr.net/npm/semantic-ui@2/dist/semantic.min.css"
                />
                <h1>众筹</h1>
                <p>当前账户：{this.state.currentAccount}</p>
                <TabCenter/>
            </div>
        );
    }
}

export default App;
