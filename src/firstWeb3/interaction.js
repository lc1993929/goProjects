let instance = require('./instance');

// console.log(instance.options.address);
// const account = '0xC8D64fdCA7DE05204b19cA62151fC4cd50Bcd106';

/*instance.methods.getMessage().call().then(value => {
    console.log(value)
    instance.methods.setMessage('1234').send({from: account})
        .then(function (receipt) {
            console.log(receipt);
            instance.methods.getMessage().call().then(value => console.log(value));
        });
});*/


let test = async () => {
    //获取账户
    let accounts = await web3.eth.getAccounts();
    console.log(accounts);
    let account = accounts[0];

    let msg1 = await instance.methods.getMessage().call();
    console.log(msg1);
    let setValue = await instance.methods.setMessage('test2').send({from: account});
    console.log(setValue);
    let msg2 = await instance.methods.getMessage().call();
    console.log(msg2);
}

test();

