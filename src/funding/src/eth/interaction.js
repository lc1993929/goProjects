import {fundingFactoryInstance, getFundingInstance} from "./instance";
import web3 from "../utils/initWeb3";


function fundingAddressToDetail(fundingAddresses) {
    return fundingAddresses.map(async function (fundingAddress) {
        let fundingInstance = getFundingInstance(fundingAddress);
        let manager = await fundingInstance.methods.manager().call();
        let projectName = await fundingInstance.methods.projectName().call();
        let targetMoney = await fundingInstance.methods.targetMoney().call();
        let supportMoney = await fundingInstance.methods.supportMoney().call();
        let remainTime = await fundingInstance.methods.getRemainTime().call();
        let balance = await fundingInstance.methods.getBalance().call();
        let supportersCount = await fundingInstance.methods.getSupportersCount().call();
        return {fundingAddress, manager, projectName, targetMoney, supportMoney, remainTime, balance, supportersCount};
    });
}

let getAllFundings = async () => {
    let fundingAddresses = await fundingFactoryInstance.methods.getAllFundings().call();
    let details = fundingAddressToDetail(fundingAddresses);

    return await Promise.all(details);

}


let getCreatorFundings = async () => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingAddresses = await fundingFactoryInstance.methods.getCreatorFundings().call({
        from: currentAccount
    });
    let details = fundingAddressToDetail(fundingAddresses);

    return await Promise.all(details);

}

let getSupporterFundings = async () => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingAddresses = await fundingFactoryInstance.methods.getSupporterFundings().call({
        from: currentAccount
    });
    let details = fundingAddressToDetail(fundingAddresses);

    return await Promise.all(details);

}


let createFunding = async (projectName, targetMoney, supportMoney, duration) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    return await fundingFactoryInstance.methods.createFunding(projectName, targetMoney, supportMoney, duration).send({
        from: currentAccount,
        gas: 2000000
    });
}

let support = async (fundingAddress, supportMoney) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingInstance = getFundingInstance(fundingAddress);
    return await fundingInstance.methods.support().send({
        from: currentAccount,
        gas: 2000000,
        value: supportMoney
    });
}

let createRequest = async (purpose, cost, shopAddress, fundingAddress) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingInstance = getFundingInstance(fundingAddress);
    return await fundingInstance.methods.createRequest(purpose, cost, shopAddress).send({
        from: currentAccount,
        gas: 2000000,
    });
}

let listRequestsByFundingAddress = async (fundingAddress) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingInstance = getFundingInstance(fundingAddress);

    let requestCount = await fundingInstance.methods.getRequestsCount().call({from: currentAccount});

    let requestDetails = [];
    for (let i = 0; i < requestCount; i++) {
        let requestDetail = await fundingInstance.methods.getRequestDetailByIndex(i).call({from: currentAccount});
        requestDetails.push(requestDetail);
    }
    return requestDetails;
}

let getSupportersCount = async (fundingAddress) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingInstance = getFundingInstance(fundingAddress);

    return await fundingInstance.methods.getSupportersCount().call({from: currentAccount});
}

let approveRequest = async (fundingAddress, index) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingInstance = getFundingInstance(fundingAddress);

    return await fundingInstance.methods.approveRequest(index).send({
        from: currentAccount,
        gas: 2000000
    });
}

let finalizeRequest = async (fundingAddress, index) => {
    let accounts = await web3.eth.getAccounts();
    let currentAccount = accounts[0];
    let fundingInstance = getFundingInstance(fundingAddress);

    return await fundingInstance.methods.finalizeRequest(index).send({
        from: currentAccount,
        gas: 2000000
    });
}


export {
    getAllFundings,
    getCreatorFundings,
    getSupporterFundings,
    createFunding,
    support,
    createRequest,
    listRequestsByFundingAddress,
    getSupportersCount,
    approveRequest,
    finalizeRequest
};