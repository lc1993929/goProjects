var HelloWorld = artifacts.require("HelloWorld");

module.exports = function(deployer) {
    // 任务就是 部署迁移合约
    deployer.deploy(HelloWorld);
};