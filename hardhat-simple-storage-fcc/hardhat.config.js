require("@nomiclabs/hardhat-waffle");
require("dotenv").config()
require("@nomicfoundation/hardhat-verify");
require("./tasks/block-number")
require("hardhat-gas-reporter")
//require("solidity-converage") 依赖版本冲突,
// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL
const PRIVATE_KEY = process.env.PRIVATE_KEY
const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY
const COIN_API_KEY = process.env.COIN_API_KEY
module.exports = {
  defaultNetwork:"hardhat",
  networks:{
    sepolia:{
      url:SEPOLIA_RPC_URL,
      accounts:[PRIVATE_KEY],
      chainId:11155111
    },
    location:{
      url:"http://127.0.0.1:8545/",
      chainId:31337
    }
  },
  solidity: "0.8.4",
  etherscan: {
    apiKey: {
      sepolia: ETHERSCAN_API_KEY
    }
  },
  //尝试调整配置,表格还是无法打印在控制台
  gasReporter:{
    enabled:true,
    noColors:true,
    currency:"USD",
    coinmarketcap:COIN_API_KEY,
    reportFormat: "markdown",
    outputFile: "gasReport.md",
  },
  abiBin: {
    // 指定合约文件的路径
    contractsPath: "./contracts",
  },
};

//报错TypeError: Class extends value undefined is not a constructor or null 解决办法降版本到1.0.7