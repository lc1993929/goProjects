import web3 from '../utils/initWeb3';

let fundingFactoryAbiFile = require('../contracts/artifacts/FundingFactory_metadata.json');
let fundingAbiFile = require('../contracts/artifacts/Funding_metadata.json');

const address = '0x1a9250635708516Ad384bAAB6Fc3aE09F974D6b0';
let fundingFactoryAbi = fundingFactoryAbiFile.output.abi;
let fundingAbi = fundingAbiFile.output.abi;

let fundingFactoryInstance = new web3.eth.Contract(fundingFactoryAbi, address);
let getFundingInstance = (address) => {
    return new web3.eth.Contract(fundingAbi, address);
}
// console.log(fundingFactoryInstance.options.address);

export {
    fundingFactoryInstance, getFundingInstance
}