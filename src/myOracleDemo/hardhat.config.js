require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: {
        compilers: [
            {
                version: "0.5.0"
            },
            {
                version: "0.8.0"
            }
        ]
    },
    networks: {
        ganache: {
            url: "HTTP://127.0.0.1:7545"
        },
        /*        goerli: {
                    url: "https://goerli.infura.io/v3/9b08343587a3477f9be3a21525d7baaf",
                    // url: "https://goerli.etherscan.io",
                    accounts: ["privateKey"]
                }*/
    }
};
